package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
)

const (
	// defaultBlockTime is the approximate waiting time for Farcaster Hub to generate new events.
	defaultBlockTime = 3 * time.Second
)

var _ engine.DataSource = (*dataSource)(nil)

type dataSource struct {
	config                  *config.Module
	option                  *Option
	farcasterClient         farcaster.Client
	databaseClient          database.Client
	state                   State
	pendingState            State
	startFarcasterTimestamp uint32
}

func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	zap.L().Info("Starting Farcaster data source")

	if err := s.initialize(); err != nil {
		zap.L().Error("Failed to initialize Farcaster data source", zap.Error(err))
		errorChan <- fmt.Errorf("initialize dataSource: %w", err)

		return
	}

	zap.L().Debug("Starting historical casts polling")
	// poll historical casts
	go func() {
		if err := retryOperation(ctx, func(ctx context.Context) error {
			return s.pollCasts(ctx, tasksChan)
		}); err != nil {
			zap.L().Error("Failed to poll historical casts", zap.Error(err))
			errorChan <- err
		} else {
			zap.L().Debug("Completed polling historical casts")
			return
		}
	}()

	zap.L().Debug("Starting historical reactions polling")
	// poll historical reactions
	go func() {
		if err := retryOperation(ctx, func(ctx context.Context) error {
			return s.pollReactions(ctx, tasksChan)
		}); err != nil {
			zap.L().Error("Failed to poll historical reactions", zap.Error(err))
			errorChan <- err
		} else {
			zap.L().Debug("Completed polling historical reactions")
			return
		}
	}()

	zap.L().Debug("Starting latest events polling")
	// poll latest events
	go func() {
		if err := retryOperation(ctx, func(ctx context.Context) error {
			return s.pollEvents(ctx, tasksChan)
		}); err != nil {
			zap.L().Error("Failed to poll latest events", zap.Error(err))
			errorChan <- err
		}
	}()

	zap.L().Info("Successfully started Farcaster data source")
}

func (s *dataSource) initialize() (err error) {
	client, err := farcaster.NewClient(s.config.Endpoint.URL, farcaster.WithAPIKey(s.option.APIKey))
	if err != nil {
		return fmt.Errorf("create farcaster client: %w", err)
	}

	s.farcasterClient = client

	return nil
}

// pollCasts polls casts from the Farcaster Hub.
// It will poll casts by fid from the maximum fid to the minimal fid (1).
func (s *dataSource) pollCasts(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	// Check if backfill of casts is complete.
	if s.state.CastsBackfill {
		zap.L().Debug("Casts backfill is already complete")
		return nil
	}

	// If fid is 0 and backfill is not complete, fetch the maximum fid from the Farcaster Hub.
	if s.state.CastsFid == 0 && !s.state.CastsBackfill {
		zap.L().Info("Fetching maximum fid from Farcaster Hub")

		fidsResponse, err := s.farcasterClient.GetFids(ctx, true, lo.ToPtr(1))

		if err != nil {
			zap.L().Error("Failed to fetch maximum fid", zap.Error(err))

			return fmt.Errorf("failed to fetch farcaster max fid: %w", err)
		}

		s.pendingState.CastsFid = fidsResponse.Fids[0]
		zap.L().Info("Successfully fetched maximum fid", zap.Uint64("fid", fidsResponse.Fids[0]))
	}

	// Poll casts by fid until backfill is complete.
	for !s.state.CastsBackfill {
		zap.L().Debug("Polling casts for fid", zap.Uint64("fid", s.pendingState.CastsFid))

		if err := s.pollCastsByFid(ctx, lo.ToPtr(int64(s.pendingState.CastsFid)), "", tasksChan); err != nil {
			zap.L().Error("Failed to poll casts for fid", zap.Uint64("fid", s.pendingState.CastsFid), zap.Error(err))

			return err
		}

		// Update state with pending state and decrement pending fid.
		s.state = s.pendingState
		s.pendingState.CastsFid--

		// If pending fid is 0, mark backfill as complete.
		if s.pendingState.CastsFid == 0 {
			zap.L().Info("Completed casts backfill")

			s.pendingState.CastsBackfill = true
			s.state = s.pendingState
		}
	}

	return nil
}

// pollCastsByFid polls casts by fid from the Farcaster Hub and send tasks to tasksChan.
func (s *dataSource) pollCastsByFid(ctx context.Context, fid *int64, pageToken string, tasksChan chan<- *engine.Tasks) error {
	for {
		// Fetch casts by fid.
		zap.L().Debug("Fetching casts by fid", zap.Int64("fid", *fid), zap.String("pageToken", pageToken))
		castsByFidResponse, err := s.farcasterClient.GetCastsByFid(ctx, fid, true, nil, pageToken)

		if err != nil {
			zap.L().Error("Failed to fetch casts by fid", zap.Int64("fid", *fid), zap.Error(err))
			return err
		}

		messages := lo.Filter(castsByFidResponse.Messages, func(message farcaster.Message, _ int) bool {
			return message.Data.Timestamp >= s.startFarcasterTimestamp
		})

		zap.L().Debug("Filtered messages by timestamp",
			zap.Int("total", len(castsByFidResponse.Messages)),
			zap.Int("filtered", len(messages)))

		// Build tasks from the fetched casts and send them to the tasks channel.
		tasks := s.buildFarcasterMessageTasks(ctx, messages)

		tasksChan <- tasks

		// If the fetched casts do not have a next page token
		// or the number of messages is less than the number of fetched messages
		if castsByFidResponse.NextPageToken == "" || len(messages) < len(castsByFidResponse.Messages) {
			zap.L().Debug("Completed fetching casts for fid", zap.Int64("fid", *fid))
			return nil
		}
		// Update the page token for the next iteration
		pageToken = castsByFidResponse.NextPageToken
	}
}

// pollReactions polls reactions from the Farcaster Hub.
// It will poll reactions by fid from the maximum fid to the minimal fid (1).
func (s *dataSource) pollReactions(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	// Check if backfill of reactions is complete.
	if s.state.ReactionsBackfill {
		zap.L().Debug("Reactions backfill is already complete")
		return nil
	}

	// If fid is 0 and backfill is not complete, fetch the maximum fid from the Farcaster Hub.
	if s.state.ReactionsFid == 0 && !s.state.ReactionsBackfill {
		zap.L().Info("Fetching maximum fid from Farcaster Hub")

		fidsResponse, err := s.farcasterClient.GetFids(ctx, true, lo.ToPtr(1))

		if err != nil {
			zap.L().Error("Failed to fetch maximum fid", zap.Error(err))

			return fmt.Errorf("failed to fetch farcaster max fid: %w", err)
		}

		s.pendingState.ReactionsFid = fidsResponse.Fids[0]
		zap.L().Info("Successfully fetched maximum fid", zap.Uint64("fid", fidsResponse.Fids[0]))
	}

	// Poll reactions by fid until backfill is complete.
	for !s.state.ReactionsBackfill {
		zap.L().Debug("Polling reactions for fid", zap.Uint64("fid", s.pendingState.ReactionsFid))

		if err := s.pollReactionsByFid(ctx, lo.ToPtr(int64(s.pendingState.ReactionsFid)), "", tasksChan); err != nil {
			zap.L().Error("Failed to poll reactions for fid", zap.Uint64("fid", s.pendingState.ReactionsFid), zap.Error(err))

			return err
		}

		// Update state with pending state and decrement pending fid.
		s.state = s.pendingState
		s.pendingState.ReactionsFid--

		// If pending fid is 0, mark backfill as complete.
		if s.pendingState.ReactionsFid == 0 {
			zap.L().Info("Completed reactions backfill")

			s.pendingState.ReactionsBackfill = true
			s.state = s.pendingState
		}
	}

	return nil
}

// pollReactionsByFid polls reactions by fid from the Farcaster Hub and send tasks to tasksChan.
func (s *dataSource) pollReactionsByFid(ctx context.Context, fid *int64, pageToken string, tasksChan chan<- *engine.Tasks) error {
	for {
		// Fetch reactions by fid.
		zap.L().Debug("Fetching reactions by fid", zap.Int64("fid", *fid), zap.String("pageToken", pageToken))
		reactionsByFidResponse, err := s.farcasterClient.GetReactionsByFid(ctx, fid, true, nil, pageToken, farcaster.ReactionTypeRecast.String())

		if err != nil {
			zap.L().Error("Failed to fetch reactions by fid", zap.Int64("fid", *fid), zap.Error(err))
			return err
		}

		messages := lo.Filter(reactionsByFidResponse.Messages, func(message farcaster.Message, _ int) bool {
			return message.Data.Timestamp >= s.startFarcasterTimestamp
		})

		zap.L().Debug("Filtered messages by timestamp",
			zap.Int("total", len(reactionsByFidResponse.Messages)),
			zap.Int("filtered", len(messages)))

		// Build tasks from the fetched reactions and send them to the tasks channel.
		tasks := s.buildFarcasterMessageTasks(ctx, messages)

		tasksChan <- tasks
		// If the fetched reactions do not have a next page token
		// or the number of messages is less than the number of fetched messages
		if reactionsByFidResponse.NextPageToken == "" || len(messages) < len(reactionsByFidResponse.Messages) {
			zap.L().Debug("Completed fetching reactions for fid", zap.Int64("fid", *fid))
			return nil
		}
		// Update the page token for the next iteration
		pageToken = reactionsByFidResponse.NextPageToken
	}
}

// buildFarcasterMessageTasks builds tasks from a slice of Farcaster messages.
func (s *dataSource) buildFarcasterMessageTasks(ctx context.Context, messages []farcaster.Message) *engine.Tasks {
	var tasks engine.Tasks

	if len(messages) == 0 {
		zap.L().Debug("No messages to build tasks")
		return &tasks
	}

	zap.L().Debug("Building tasks from messages",
		zap.Int("messages.count", len(messages)))

	resultPool := pool.NewWithResults[*Task]().WithMaxGoroutines(lo.Ternary(len(messages) < 20*runtime.NumCPU(), len(messages), 20*runtime.NumCPU()))

	for _, message := range messages {
		message := message

		resultPool.Go(func() *Task {
			switch message.Data.Type {
			case farcaster.MessageTypeCastAdd.String():
				if err := s.fillCastParams(ctx, &message); err != nil {
					zap.L().Warn("Failed to fill cast parameters",
						zap.Uint64("fid", message.Data.Fid),
						zap.String("hash", message.Hash),
						zap.Error(err))

					return nil
				}
			case farcaster.MessageTypeReactionAdd.String():
				if err := s.fillReactionParams(ctx, &message); err != nil {
					zap.L().Warn("Failed to fill reaction parameters",
						zap.Uint64("fid", message.Data.Fid),
						zap.String("hash", message.Hash),
						zap.Error(err))

					return nil
				}
			}

			return &Task{
				Network: s.Network(),
				Message: message,
			}
		})
	}

	for _, task := range resultPool.Wait() {
		if task != nil {
			tasks.Tasks = append(tasks.Tasks, task)
		}
	}

	zap.L().Debug("Successfully built tasks from messages",
		zap.Int("tasks.count", len(tasks.Tasks)))

	return &tasks
}

// pollEvents polls events from the Farcaster Hub.
// It will poll events from the last event id to the latest event id.
func (s *dataSource) pollEvents(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	// Set the cursor to the current event ID in the state.
	cursor := s.state.EventID

	if cursor == 0 {
		// If the cursor is 0, fetch the latest event ID from the Farcaster Hub.
		cursor = farcaster.ConvertTimestampMilliToEventID(time.Now().Add(-10 * time.Second).UnixMilli())
		zap.L().Debug("Starting event polling from current timestamp",
			zap.Uint64("event.id", cursor))
	}

	for {
		// Fetch events from the Farcaster Hub using the cursor.
		zap.L().Info("Fetching events from Farcaster Hub",
			zap.Uint64("event.from.id", cursor))

		eventsResponse, err := s.farcasterClient.GetEvents(ctx, lo.ToPtr(int64(cursor)))

		if err != nil || eventsResponse == nil {
			zap.L().Error("Failed to fetch events from Farcaster Hub",
				zap.Uint64("event.from.id", cursor),
				zap.Error(err))

			time.Sleep(defaultBlockTime)

			continue
		}

		// If the fetched events are empty, log an info message, wait for a default block time, and continue to the next iteration.
		if len(eventsResponse.Events) == 0 {
			zap.L().Info("No new events found, waiting for next poll",
				zap.Uint64("event.from.id", cursor),
				zap.Duration("block.time", defaultBlockTime))

			time.Sleep(defaultBlockTime)

			continue
		}

		zap.L().Debug("Processing fetched events",
			zap.Int("events.count", len(eventsResponse.Events)))

		tasks := s.buildFarcasterEventTasks(ctx, eventsResponse.Events, tasksChan)

		tasksChan <- tasks

		s.state = s.pendingState
		s.pendingState.EventID = eventsResponse.NextPageEventID

		cursor = eventsResponse.NextPageEventID

		zap.L().Debug("Successfully processed events batch",
			zap.Uint64("next.event.id", cursor))
	}
}

// buildFarcasterEventTasks filter different types of events and build tasks from them.
func (s *dataSource) buildFarcasterEventTasks(ctx context.Context, events []farcaster.HubEvent, tasksChan chan<- *engine.Tasks) *engine.Tasks {
	var tasks engine.Tasks

	if len(events) == 0 {
		zap.L().Debug("No events to process")
		return &tasks
	}

	zap.L().Info("Building tasks from events", zap.Int("events.count", len(events)))

	resultPool := pool.NewWithResults[*Task]().WithMaxGoroutines(lo.Ternary(len(events) < 20*runtime.NumCPU(), len(events), 20*runtime.NumCPU()))

	for _, event := range events {
		event := event

		resultPool.Go(func() *Task {
			if event.Type != farcaster.HubEventTypeMergeMessage.String() {
				zap.L().Debug("Skipping non-merge message event", zap.String("event.type", event.Type))
				return nil
			}

			message := event.MergeMessageBody.Message
			switch message.Data.Type {
			case farcaster.MessageTypeCastAdd.String():
				if err := s.fillCastParams(ctx, &message); err != nil {
					zap.L().Warn("Failed to fill cast parameters",
						zap.Uint64("fid", message.Data.Fid),
						zap.String("hash", message.Hash),
						zap.Error(err))

					return nil
				}

				zap.L().Debug("Successfully filled cast parameters", zap.String("hash", message.Hash))

			case farcaster.MessageTypeReactionAdd.String():
				if message.Data.ReactionBody.Type == farcaster.ReactionTypeRecast.String() {
					if err := s.fillReactionParams(ctx, &message); err != nil {
						zap.L().Warn("Failed to fill reaction parameters",
							zap.Uint64("fid", message.Data.Fid),
							zap.String("hash", message.Hash),
							zap.Error(err))

						return nil
					}

					zap.L().Debug("Successfully filled recast parameters", zap.String("hash", message.Hash))
				} else {
					zap.L().Debug("Skipping non-recast reaction", zap.String("reaction.type", message.Data.ReactionBody.Type))
					return nil
				}

			case farcaster.MessageTypeVerificationRemove.String(),
				farcaster.MessageTypeVerificationAddEthAddress.String(),
				farcaster.MessageTypeUserDataAdd.String(),
				farcaster.MessageTypeUsernameProof.String():
				fid := int64(message.Data.Fid)

				zap.L().Info("Processing verification/user data message",
					zap.Int64("fid", fid),
					zap.String("message.type", message.Data.Type))

				_, _ = s.updateProfileByFid(ctx, &fid)

				if message.Data.Type == farcaster.MessageTypeVerificationAddEthAddress.String() {
					zap.L().Info("Polling casts and reactions for new ETH address verification",
						zap.Int64("fid", fid))

					_ = s.pollCastsByFid(ctx, &fid, "", tasksChan)
					_ = s.pollReactionsByFid(ctx, &fid, "", tasksChan)
				}

				return nil
			default:
				zap.L().Debug("Skipping unsupported message type", zap.String("type", message.Data.Type))
				return nil
			}

			return &Task{
				Network: s.Network(),
				Message: message,
			}
		})
	}

	results := resultPool.Wait()
	for _, task := range results {
		if task != nil {
			tasks.Tasks = append(tasks.Tasks, task)
		}
	}

	zap.L().Info("Successfully built tasks from events",
		zap.Int("total.events", len(events)),
		zap.Int("processed.tasks", len(tasks.Tasks)))

	return &tasks
}

// updateProfileByFid update profile by fid.
// It will fetch the username and custody address by fid and update the profile in the database.
func (s *dataSource) updateProfileByFid(ctx context.Context, fid *int64) (*model.Profile, error) {
	var username string

	zap.L().Debug("Fetching user data by fid", zap.Int64("fid", *fid))
	// owner username(handle)
	userData, err := s.farcasterClient.GetUserDataByFidAndType(ctx, fid, farcaster.UserDataTypeUsername.String())

	if err != nil {
		zap.L().Error("Failed to fetch user data by fid", zap.Error(err), zap.Int64("fid", *fid))
	} else if userData.Data.UserDataBody != nil {
		username = userData.Data.UserDataBody.Value
		zap.L().Debug("Found username for fid", zap.String("username", username), zap.Int64("fid", *fid))
	}

	// custody address
	zap.L().Debug("Fetching custody address", zap.Int64("fid", *fid))
	custodyAddress, username, err := s.getCustodyAddress(ctx, fid, username)

	if err != nil {
		zap.L().Error("Failed to get custody address", zap.Error(err), zap.Int64("fid", *fid))
		return nil, err
	}

	// eth addresses
	zap.L().Debug("Fetching ETH addresses", zap.Int64("fid", *fid))
	ethAddresses, err := s.getEthAddresses(ctx, fid)

	if err != nil {
		zap.L().Error("Failed to get ETH addresses", zap.Error(err), zap.Int64("fid", *fid))

		return nil, err
	}

	profile := &model.Profile{
		Fid:            *fid,
		Username:       username,
		CustodyAddress: common.HexToAddress(custodyAddress).String(),
		EthAddresses:   ethAddresses,
	}

	zap.L().Debug("Saving profile to database", zap.Int64("fid", *fid))

	if err = s.databaseClient.SaveDatasetFarcasterProfile(ctx, profile); err != nil {
		zap.L().Error("Failed to save profile", zap.Error(err), zap.Int64("fid", *fid))
		return nil, err
	}

	zap.L().Info("Successfully updated profile", zap.Int64("fid", *fid))

	return profile, nil
}

// getCustodyAddress get custody address by fid.
func (s *dataSource) getCustodyAddress(ctx context.Context, fid *int64, username string) (string, string, error) {
	zap.L().Debug("Fetching username proofs", zap.Int64("fid", *fid))
	userProofs, err := s.farcasterClient.GetUserNameProofsByFid(ctx, fid)

	if err != nil {
		zap.L().Error("Failed to fetch username proofs", zap.Error(err), zap.Int64("fid", *fid))
		return "", "", fmt.Errorf("fetch custody address by fid error: %w,%d", err, fid)
	}

	var custodyAddress string

	for _, proof := range userProofs.Proofs {
		if proof.Type == farcaster.UsernameTypeFname.String() {
			custodyAddress = proof.Owner

			if username == "" {
				username = proof.Name
			}

			zap.L().Debug("Found custody address from proofs",
				zap.String("address", custodyAddress),
				zap.String("username", username))

			break
		}
	}

	if custodyAddress == "" {
		zap.L().Debug("Fetching name proof by username", zap.String("username", username))
		nameProof, err := s.farcasterClient.GetUserNameProofByName(ctx, username)

		if err != nil || nameProof == nil {
			zap.L().Error("Failed to fetch name proof", zap.Error(err), zap.String("username", username))

			return "", "", fmt.Errorf("fetch custody address by name error: %w,%d", err, fid)
		}

		custodyAddress = nameProof.Owner
		zap.L().Debug("Found custody address from name proof", zap.String("address", custodyAddress))
	}

	return custodyAddress, username, nil
}

// getEthAddresses get eth addresses by fid.
func (s *dataSource) getEthAddresses(ctx context.Context, fid *int64) ([]string, error) {
	zap.L().Debug("Fetching verifications", zap.Int64("fid", *fid))
	verifications, err := s.farcasterClient.GetVerificationsByFid(ctx, fid, "")

	if err != nil {
		zap.L().Error("Failed to fetch verifications", zap.Error(err), zap.Int64("fid", *fid))
		return nil, fmt.Errorf("failed to fetch eth address for fid %d: %w", fid, err)
	}

	var ethAddresses []string

	for _, msg := range verifications.Messages {
		if msg.Data.Type == farcaster.MessageTypeVerificationAddEthAddress.String() && msg.Data.VerificationAddEthAddressBody != nil && msg.Data.VerificationAddEthAddressBody.Protocol == farcaster.ProtocolEthereum.String() {
			address := common.HexToAddress(msg.Data.VerificationAddEthAddressBody.Address).String()
			ethAddresses = append(ethAddresses, address)
			zap.L().Debug("Found ETH address", zap.String("address", address))
		}
	}

	return ethAddresses, nil
}

// getProfileByFid get profile by fid.
// It will fetch the profile by fid from the database.
// If the profile is not found or the username, custody address, and eth addresses are empty, it will update the profile.
func (s *dataSource) getProfileByFid(ctx context.Context, fid *int64) (*model.Profile, error) {
	var (
		profile *model.Profile
		err     error
	)

	zap.L().Debug("Loading profile from database", zap.Int64("fid", *fid))
	profile, err = s.databaseClient.LoadDatasetFarcasterProfile(ctx, *fid)

	if err != nil {
		zap.L().Error("Failed to load profile", zap.Error(err), zap.Int64("fid", *fid))
		return nil, err
	}

	if profile != nil && profile.Username == "" && profile.CustodyAddress == "" && len(profile.EthAddresses) == 0 {
		zap.L().Info("Profile incomplete, updating profile data", zap.Int64("fid", *fid))
		profile, err = s.updateProfileByFid(ctx, fid)

		if err != nil {
			zap.L().Error("Failed to update profile", zap.Error(err), zap.Int64("fid", *fid))
			return nil, fmt.Errorf("failed to update farcaster profile for fid %d: %w", fid, err)
		}
	}

	return profile, nil
}

// fillMentionsUsernames transfer mentions fid to username.
func (s *dataSource) fillMentionsUsernames(ctx context.Context, message *farcaster.Message) error {
	zap.L().Debug("Starting to fill mentions usernames",
		zap.Int("mentions_count", len(message.Data.CastAddBody.Mentions)))

	message.Data.CastAddBody.MentionsUsernames = make([]string, len(message.Data.CastAddBody.Mentions))
	for i, fid := range message.Data.CastAddBody.Mentions {
		fid := int64(fid)
		profile, err := s.getProfileByFid(ctx, &fid)

		if err != nil {
			zap.L().Warn("Failed to fetch profile for mention",
				zap.Int64("fid", fid),
				zap.Error(err))

			continue
		}

		message.Data.CastAddBody.MentionsUsernames[i] = profile.Username
		zap.L().Debug("Successfully filled mention username",
			zap.Int64("fid", fid),
			zap.String("username", profile.Username))
	}

	return nil
}

// fillProfile fill profile in each message.
func (s *dataSource) fillProfile(ctx context.Context, message *farcaster.Message) error {
	fid := int64(message.Data.Fid)
	zap.L().Debug("Fetching profile for message", zap.Int64("fid", fid))

	profile, err := s.getProfileByFid(ctx, &fid)

	if err != nil {
		zap.L().Error("Failed to fetch profile for message",
			zap.Int64("fid", fid),
			zap.Error(err))

		return fmt.Errorf("failed to fetch farcaster profile for fid %d: %w", fid, err)
	}

	message.Data.Profile = profile

	zap.L().Debug("Successfully filled profile for message", zap.Int64("fid", fid))

	return nil
}

// fillCastParams fill params in cast add message.
func (s *dataSource) fillCastParams(ctx context.Context, message *farcaster.Message) error {
	if message.Data.CastAddBody.ParentCastID != nil {
		targetFid := int64(message.Data.CastAddBody.ParentCastID.Fid)
		zap.L().Debug("Fetching parent cast",
			zap.Int64("target_fid", targetFid),
			zap.String("hash", message.Data.CastAddBody.ParentCastID.Hash))

		targetMessage, err := s.farcasterClient.GetCastByFidAndHash(ctx, &targetFid, message.Data.CastAddBody.ParentCastID.Hash)

		if err != nil {
			zap.L().Error("Failed to fetch parent cast",
				zap.Int64("target_fid", targetFid),
				zap.Error(err))

			return fmt.Errorf("failed to fetch target cast for target fid %d: %w", targetFid, err)
		}

		if targetMessage == nil {
			zap.L().Error("Parent cast not found", zap.Int64("target_fid", targetFid))
			return fmt.Errorf("not found target cast for target fid %d", targetFid)
		}

		if err = s.fillMentionsUsernames(ctx, targetMessage); err != nil {
			return err
		}

		message.Data.CastAddBody.ParentCast = targetMessage

		if err = s.fillProfile(ctx, message.Data.CastAddBody.ParentCast); err != nil {
			return err
		}
	}

	if err := s.fillProfile(ctx, message); err != nil {
		return err
	}

	return s.fillMentionsUsernames(ctx, message)
}

// fillReactionParams fill params in recast message.
func (s *dataSource) fillReactionParams(ctx context.Context, message *farcaster.Message) error {
	if message.Data.ReactionBody.Type == farcaster.ReactionTypeRecast.String() {
		targetFid := int64(message.Data.ReactionBody.TargetCastID.Fid)
		zap.L().Debug("Fetching target cast for reaction",
			zap.Int64("target_fid", targetFid),
			zap.String("hash", message.Data.ReactionBody.TargetCastID.Hash))

		targetMessage, err := s.farcasterClient.GetCastByFidAndHash(ctx, &targetFid, message.Data.ReactionBody.TargetCastID.Hash)

		if err != nil {
			zap.L().Error("Failed to fetch target cast for reaction",
				zap.Int64("target_fid", targetFid),
				zap.Error(err))

			return fmt.Errorf("failed to fetch target reaction for target fid %d: %w", targetFid, err)
		}

		if targetMessage == nil {
			zap.L().Error("Target cast for reaction not found", zap.Int64("target_fid", targetFid))
			return fmt.Errorf("not found target reaction for target fid %d", targetFid)
		}

		if err = s.fillMentionsUsernames(ctx, targetMessage); err != nil {
			return err
		}

		message.Data.ReactionBody.TargetCast = targetMessage

		if err = s.fillProfile(ctx, message.Data.ReactionBody.TargetCast); err != nil {
			return err
		}
	}

	return s.fillProfile(ctx, message)
}

func retryOperation(ctx context.Context, operation func(ctx context.Context) error) error {
	return retry.Do(
		func() error {
			return operation(ctx)
		},
		retry.Attempts(0),
		retry.Delay(1*time.Second),
		retry.DelayType(retry.BackOffDelay),
		retry.OnRetry(func(n uint, err error) {
			zap.L().Warn("Retrying Farcaster data source operation",
				zap.Uint("attempt", n),
				zap.Error(err))
		}),
	)
}

func NewSource(config *config.Module, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.DataSource, error) {
	zap.L().Info("Creating new Farcaster data source")

	var (
		state State
		err   error
	)

	// Initialize state from checkpoint.
	if checkpoint != nil {
		zap.L().Debug("Initializing state from checkpoint")

		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			zap.L().Error("Failed to unmarshal checkpoint state", zap.Error(err))
			return nil, err
		}
	}

	instance := dataSource{
		databaseClient: databaseClient,
		config:         config,
		state:          state,
		pendingState:   state, // Default pending state is equal to current state.
	}

	if instance.option, err = NewOption(config.Network, config.Parameters); err != nil {
		zap.L().Error("Failed to parse config options", zap.Error(err))
		return nil, fmt.Errorf("parse config: %w", err)
	}

	zap.L().Info("Applied Farcaster options", zap.Any("option", instance.option))

	if instance.option.TimestampStart != nil {
		zap.L().Debug("Setting start timestamp",
			zap.Int64("raw_timestamp", instance.option.TimestampStart.Int64()))

		instance.startFarcasterTimestamp = farcaster.CovertTimestampToFarcasterTime(instance.option.TimestampStart.Int64())
	}

	zap.L().Info("Successfully created Farcaster data source")

	return &instance, nil
}
