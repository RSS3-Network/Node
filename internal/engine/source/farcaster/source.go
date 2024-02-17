package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	// defaultBlockTime is the approximate waiting time for Farcaster Hub to generate new events.
	defaultBlockTime = 3 * time.Second
)

var _ engine.Source = (*source)(nil)

type source struct {
	config          *config.Module
	option          *Option
	farcasterClient farcaster.Client
	databaseClient  database.Client
	state           State
	pendingState    State
}

func (s *source) Network() filter.Network {
	return s.config.Network
}

func (s *source) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *source) Start(ctx context.Context, tasksChan chan<- []engine.Task, errorChan chan<- error) {
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("initialize source: %w", err)

		return
	}

	// poll historical casts
	go func() {
		err := s.pollCasts(ctx, tasksChan)
		if err == nil {
			return
		}

		errorChan <- err
	}()

	// poll historical reactions
	go func() {
		if err := s.pollReactions(ctx, tasksChan); err != nil {
			errorChan <- err
		} else {
			return
		}
	}()

	// poll latest events
	go func() {
		errorChan <- s.pollEvents(ctx, tasksChan)
	}()
}

func (s *source) initialize() (err error) {
	client, err := farcaster.NewClient(s.config.Endpoint, farcaster.WithAPIKey(s.option.APIKey))
	if err != nil {
		return fmt.Errorf("create farcaster client: %w", err)
	}

	s.farcasterClient = client

	return nil
}

func (s *source) pollCasts(ctx context.Context, tasksChan chan<- []engine.Task) error {
	if s.state.CastsBackfill {
		return nil
	}

	if s.state.CastsFid == 0 && !s.state.CastsBackfill {
		fidsResponse, err := s.farcasterClient.GetFids(ctx, true, lo.ToPtr(1))
		if err != nil {
			return fmt.Errorf("fetch farcaster max fid: %w", err)
		}

		s.pendingState.CastsFid = fidsResponse.Fids[0]
	}

	for !s.state.CastsBackfill {
		err := s.pollCastsByFid(ctx, lo.ToPtr(int64(s.pendingState.CastsFid)), "", tasksChan)

		if err != nil {
			return err
		}

		s.state = s.pendingState
		s.pendingState.CastsFid--

		if s.pendingState.CastsFid == 0 {
			s.pendingState.CastsBackfill = true
			s.state = s.pendingState
		}
	}

	return nil
}

func (s *source) pollCastsByFid(ctx context.Context, fid *int64, pageToken string, tasksChan chan<- []engine.Task) error {
	for {
		castsByFidResponse, err := s.farcasterClient.GetCastsByFid(ctx, fid, true, nil, pageToken)

		if err != nil {
			return err
		}

		tasks, err := s.buildFarcasterMessageTasks(ctx, castsByFidResponse.Messages)
		if err != nil {
			return fmt.Errorf("build message tasks: %w", err)
		}

		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		if castsByFidResponse.NextPageToken == "" {
			return nil
		}

		pageToken = castsByFidResponse.NextPageToken
	}
}

func (s *source) pollReactions(ctx context.Context, tasksChan chan<- []engine.Task) error {
	if s.state.ReactionsBackfill {
		return nil
	}

	if s.state.ReactionsFid == 0 && !s.state.ReactionsBackfill {
		fidsResponse, err := s.farcasterClient.GetFids(ctx, true, lo.ToPtr(1))
		if err != nil {
			return fmt.Errorf("fetch farcaster max fid: %w", err)
		}

		s.pendingState.ReactionsFid = fidsResponse.Fids[0]
	}

	for !s.state.ReactionsBackfill {
		err := s.pollReactionsByFid(ctx, lo.ToPtr(int64(s.pendingState.ReactionsFid)), "", tasksChan)

		if err != nil {
			return err
		}

		s.state = s.pendingState
		s.pendingState.ReactionsFid--

		if s.pendingState.ReactionsFid == 0 {
			s.pendingState.ReactionsBackfill = true
			s.state = s.pendingState
		}
	}

	return nil
}

func (s *source) pollReactionsByFid(ctx context.Context, fid *int64, pageToken string, tasksChan chan<- []engine.Task) error {
	for {
		reactionsByFidResponse, err := s.farcasterClient.GetReactionsByFid(ctx, fid, true, nil, pageToken, farcaster.ReactionTypeRecast.String())

		if err != nil {
			return err
		}

		tasks, err := s.buildFarcasterMessageTasks(ctx, reactionsByFidResponse.Messages)
		if err != nil {
			return fmt.Errorf("build message tasks: %w", err)
		}

		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		if reactionsByFidResponse.NextPageToken == "" {
			return nil
		}

		pageToken = reactionsByFidResponse.NextPageToken
	}
}

// buildFarcasterMessageTasks filter cast add and recast messages.
func (s *source) buildFarcasterMessageTasks(ctx context.Context, messages []farcaster.Message) ([]*Task, error) {
	tasks := make([]*Task, 0)

	for _, message := range messages {
		message := message
		switch message.Data.Type {
		case farcaster.MessageTypeCastAdd.String():
			if err := s.fillCastParams(ctx, &message); err != nil {
				zap.L().Warn("fill farcaster cast params error", zap.Uint64("fid", message.Data.Fid), zap.String("hash", message.Hash))

				continue
			}
		case farcaster.MessageTypeReactionAdd.String():
			if err := s.fillReactionParams(ctx, &message); err != nil {
				zap.L().Warn("fill farcaster reaction params error", zap.Uint64("fid", message.Data.Fid), zap.String("hash", message.Hash))

				continue
			}
		}

		tasks = append(tasks, &Task{
			Network: s.Network(),
			Message: message,
		})
	}

	return tasks, nil
}

func (s *source) pollEvents(ctx context.Context, tasksChan chan<- []engine.Task) error {
	cursor := s.state.EventID

	for {
		eventsResponse, err := s.farcasterClient.GetEvents(ctx, lo.ToPtr(int64(cursor)))

		if err != nil || eventsResponse == nil {
			zap.L().Warn("fetch farcaster events error", zap.Uint64("event.from.id", cursor))

			time.Sleep(defaultBlockTime)

			continue
		}

		if len(eventsResponse.Events) == 0 {
			zap.L().Info("wait for new events", zap.Uint64("event.from.id", cursor), zap.Duration("block.time", defaultBlockTime))

			time.Sleep(defaultBlockTime)

			continue
		}

		tasks, err := s.buildFarcasterEventTasks(ctx, eventsResponse.Events, tasksChan)
		if err != nil {
			return fmt.Errorf("build event tasks: %w", err)
		}
		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		s.state = s.pendingState
		s.pendingState.EventID = eventsResponse.NextPageEventID

		cursor = eventsResponse.NextPageEventID
	}
}

// buildFarcasterEventTasks filter cast add, recast and profile update events.
func (s *source) buildFarcasterEventTasks(ctx context.Context, events []farcaster.HubEvent, tasksChan chan<- []engine.Task) ([]*Task, error) {
	tasks := make([]*Task, 0)

	for _, event := range events {
		if event.Type == farcaster.HubEventTypeMergeMessage.String() {
			switch event.MergeMessageBody.Message.Data.Type {
			case farcaster.MessageTypeCastAdd.String():
				message := event.MergeMessageBody.Message
				if err := s.fillCastParams(ctx, &message); err != nil {
					zap.L().Warn("fill farcaster cast params error", zap.Uint64("fid", message.Data.Fid), zap.String("hash", message.Hash))

					continue
				}

				tasks = append(tasks, &Task{
					Network: s.Network(),
					Message: message,
				})
			case farcaster.MessageTypeReactionAdd.String():
				if event.MergeMessageBody.Message.Data.ReactionBody.Type == farcaster.ReactionTypeRecast.String() {
					message := event.MergeMessageBody.Message
					if err := s.fillReactionParams(ctx, &message); err != nil {
						zap.L().Warn("fill farcaster reaction params error", zap.Uint64("fid", message.Data.Fid), zap.String("hash", message.Hash))

						continue
					}

					tasks = append(tasks, &Task{
						Network: s.Network(),
						Message: message,
					})
				}
			case farcaster.MessageTypeVerificationRemove.String(),
				farcaster.MessageTypeVerificationAddEthAddress.String(),
				farcaster.MessageTypeUserDataAdd.String(),
				farcaster.MessageTypeUsernameProof.String():
				fid := int64(event.MergeMessageBody.Message.Data.Fid)

				_, _ = s.updateProfileByFid(ctx, &fid)

				if event.MergeMessageBody.Message.Data.Type == farcaster.MessageTypeVerificationAddEthAddress.String() {
					_ = s.pollCastsByFid(ctx, &fid, "", tasksChan)
					_ = s.pollReactionsByFid(ctx, &fid, "", tasksChan)
				}
			default:
				zap.L().Debug("unsupport message type", zap.String("type", event.MergeMessageBody.Message.Data.Type))

				continue
			}
		}
	}

	return tasks, nil
}

// updateProfileByFid update profile by fid.
func (s *source) updateProfileByFid(ctx context.Context, fid *int64) (*model.Profile, error) {
	// owner username(handle)
	userDataByFidAndTypeRes, err := s.farcasterClient.GetUserDataByFidAndType(ctx, fid, farcaster.UserDataTypeUsername.String())

	if err != nil {
		return nil, fmt.Errorf("fetch user data error: %w,%d", err, fid)
	}

	// custody address
	custodyAddresses, err := s.farcasterClient.GetUserNameProofsByFid(ctx, fid)

	if err != nil {
		return nil, fmt.Errorf("fetch custody address error: %w,%d", err, fid)
	}

	var custodyAddress string

	if len(custodyAddresses.Proofs) > 0 {
		addresses := lo.Filter(custodyAddresses.Proofs, func(x farcaster.UserNameProof, index int) bool {
			return x.Type == farcaster.UsernameTypeFname.String()
		})

		if len(addresses) > 0 {
			custodyAddress = addresses[0].Owner
		}
	}

	// eth addresses
	verificationsByFidRes, err := s.farcasterClient.GetVerificationsByFid(ctx, fid, "")

	if err != nil {
		return nil, fmt.Errorf("fetch eth address error: %w,%d", err, fid)
	}

	ethAddresses := lo.Map(verificationsByFidRes.Messages, func(x farcaster.Message, index int) string {
		return common.HexToAddress(x.Data.VerificationAddEthAddressBody.Address).String()
	})

	if err != nil {
		return nil, fmt.Errorf("get new farcaster profile %d: %w", fid, err)
	}

	profile := &model.Profile{
		Fid:            *fid,
		Username:       userDataByFidAndTypeRes.Data.UserDataBody.Value,
		CustodyAddress: common.HexToAddress(custodyAddress).String(),
		EthAddresses:   ethAddresses,
	}

	if err = s.databaseClient.SaveDatasetFarcasterProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

// getProfileByFid get profile by fid.
func (s *source) getProfileByFid(ctx context.Context, fid *int64) (*model.Profile, error) {
	var (
		profile *model.Profile
		err     error
	)

	profile, err = s.databaseClient.LoadDatasetFarcasterProfile(ctx, *fid)

	if err != nil {
		return nil, err
	}

	if profile != nil && profile.Username == "" && profile.CustodyAddress == "" && len(profile.EthAddresses) == 0 {
		profile, err = s.updateProfileByFid(ctx, fid)

		if err != nil {
			return nil, fmt.Errorf("update farcaster profile %d: %w", fid, err)
		}
	}

	return profile, nil
}

// fillMentionsUsernames transfer mentions fid to username.
func (s *source) fillMentionsUsernames(ctx context.Context, message *farcaster.Message) error {
	message.Data.CastAddBody.MentionsUsernames = make([]string, len(message.Data.CastAddBody.Mentions))
	for i, fid := range message.Data.CastAddBody.Mentions {
		fid := int64(fid)
		profile, err := s.getProfileByFid(ctx, &fid)

		if err != nil {
			zap.L().Warn("fetch farcaster profile error", zap.Int64("fid", fid))

			continue
		}

		message.Data.CastAddBody.MentionsUsernames[i] = profile.Username
	}

	return nil
}

// fillProfile fill profile in each message.
func (s *source) fillProfile(ctx context.Context, message *farcaster.Message) error {
	fid := int64(message.Data.Fid)
	profile, err := s.getProfileByFid(ctx, &fid)

	if err != nil {
		return fmt.Errorf("fetch farcaster profile error: %w", err)
	}

	message.Data.Profile = profile

	return nil
}

// fillCastParams fill params in cast add message.
func (s *source) fillCastParams(ctx context.Context, message *farcaster.Message) error {
	if message.Data.CastAddBody.ParentCastID != nil {
		targetFid := int64(message.Data.CastAddBody.ParentCastID.Fid)
		targetMessage, err := s.farcasterClient.GetCastByFidAndHash(ctx, &targetFid, message.Data.CastAddBody.ParentCastID.Hash)

		if err != nil {
			return err
		}

		if targetMessage == nil {
			return fmt.Errorf("target cast not found")
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
func (s *source) fillReactionParams(ctx context.Context, message *farcaster.Message) error {
	if message.Data.ReactionBody.Type == farcaster.ReactionTypeRecast.String() {
		targetFid := int64(message.Data.ReactionBody.TargetCastID.Fid)
		targetMessage, err := s.farcasterClient.GetCastByFidAndHash(ctx, &targetFid, message.Data.ReactionBody.TargetCastID.Hash)

		if err != nil {
			return err
		}

		if targetMessage == nil {
			return fmt.Errorf("target cast not found")
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

func NewSource(config *config.Module, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.Source, error) {
	var (
		state State

		err error
	)

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := source{
		databaseClient: databaseClient,
		config:         config,
		state:          state,
		pendingState:   state, // Default pending state is equal to current state.
	}

	if instance.option, err = NewOption(config.Parameters); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	zap.L().Info("apply option", zap.Any("option", instance.option))

	return &instance, nil
}
