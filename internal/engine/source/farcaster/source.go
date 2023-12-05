package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/farcaster/model"
	"github.com/naturalselectionlabs/rss3-node/provider/farcaster"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	// defaultBlockTime is the approximate waiting time for Farcaster Hub to generate new events.
	defaultBlockTime = 3 * time.Second
)

var _ engine.Source = (*source)(nil)

type source struct {
	config          *engine.Config
	farcasterClient farcaster.Client
	databaseClient  database.Client
	state           State
}

func (s *source) Chain() filter.Chain {
	return filter.ChainFarcasterMainnet
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
	client, err := farcaster.NewClient(s.config.Endpoint)
	if err != nil {
		return fmt.Errorf("create farcaster client: %w", err)
	}

	s.farcasterClient = client

	return nil
}

func (s *source) pollCasts(ctx context.Context, tasksChan chan<- []engine.Task) error {
	if s.state.CastsFid == 1 {
		return nil
	}

	if s.state.CastsFid == 0 {
		fidsResponse, err := s.farcasterClient.GetFids(ctx, true, lo.ToPtr(1))
		if err != nil {
			return fmt.Errorf("fetch farcaster max fid: %w", err)
		}

		s.state.CastsFid = fidsResponse.Fids[0] + 1
	}

	for s.state.CastsFid > 1 {
		s.state.CastsFid--
		err := s.pollCastsByFid(ctx, lo.ToPtr(int64(s.state.CastsFid)), "", tasksChan)

		if err != nil {
			return err
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
	if s.state.ReactionsFid == 1 {
		return nil
	}

	if s.state.ReactionsFid == 0 {
		fidsResponse, err := s.farcasterClient.GetFids(ctx, true, lo.ToPtr(1))
		if err != nil {
			return fmt.Errorf("fetch farcaster max fid: %w", err)
		}

		s.state.ReactionsFid = fidsResponse.Fids[0] + 1
	}

	for s.state.ReactionsFid > 1 {
		s.state.ReactionsFid--
		err := s.pollReactionsByFid(ctx, lo.ToPtr(int64(s.state.ReactionsFid)), "", tasksChan)

		if err != nil {
			return err
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

func (s *source) buildFarcasterMessageTasks(ctx context.Context, messages []farcaster.Message) ([]*Task, error) {
	tasks := make([]*Task, 0)

	for _, msg := range messages {
		msg := msg
		switch msg.Data.Type {
		case farcaster.MessageTypeCastAdd.String():
			if msg.Data.CastAddBody.ParentCastID != nil {
				targetFid := int64(msg.Data.CastAddBody.ParentCastID.Fid)
				targetMessage, err := s.farcasterClient.GetCastByFidAndHash(ctx, &targetFid, msg.Data.CastAddBody.ParentCastID.Hash)

				if err != nil {
					zap.L().Warn("fetch farcaster target cast error", zap.Int64("fid", targetFid), zap.String("hash", msg.Data.CastAddBody.ParentCastID.Hash))
				}

				_ = s.fillMentionsUsernames(ctx, targetMessage)

				msg.Data.CastAddBody.ParentCast = targetMessage
			}

			_ = s.fillMentionsUsernames(ctx, &msg)
		case farcaster.MessageTypeReactionAdd.String():
			if msg.Data.ReactionBody.Type == farcaster.ReactionTypeRecast.String() {
				targetFid := int64(msg.Data.ReactionBody.TargetCastID.Fid)
				targetMessage, err := s.farcasterClient.GetCastByFidAndHash(ctx, &targetFid, msg.Data.ReactionBody.TargetCastID.Hash)

				if err != nil {
					zap.L().Warn("fetch farcaster target cast error", zap.Int64("fid", targetFid), zap.String("hash", msg.Data.CastAddBody.ParentCastID.Hash))
				}

				_ = s.fillMentionsUsernames(ctx, targetMessage)

				msg.Data.ReactionBody.TargetCast = targetMessage
			}
		}

		tasks = append(tasks, &Task{
			Chain:   filter.ChainFarcaster(s.Chain().ID()),
			Message: msg,
		})
	}

	return tasks, nil
}

func (s *source) pollEvents(ctx context.Context, tasksChan chan<- []engine.Task) error {
	cursor := s.state.EventID

	for {
		eventsResponse, err := s.farcasterClient.GetEvents(ctx, lo.ToPtr(int64(cursor)))

		if err != nil {
			zap.L().Warn("fetch farcaster events error", zap.Uint64("event.from.id", cursor))

			time.Sleep(defaultBlockTime)
		}

		if len(eventsResponse.Events) == 0 {
			zap.L().Info("wait for new events", zap.Uint64("event.from.id", cursor), zap.Duration("block.time", defaultBlockTime))

			time.Sleep(defaultBlockTime)
		}

		tasks, err := s.buildFarcasterEventTasks(ctx, eventsResponse.Events, tasksChan)
		if err != nil {
			return fmt.Errorf("build event tasks: %w", err)
		}
		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		s.state.EventID = eventsResponse.NextPageEventID

		cursor = eventsResponse.NextPageEventID
	}
}

func (s *source) buildFarcasterEventTasks(ctx context.Context, events []farcaster.HubEvent, tasksChan chan<- []engine.Task) ([]*Task, error) {
	tasks := make([]*Task, 0)

	for _, event := range events {
		if event.Type == farcaster.HubEventTypeMergeMessage.String() {
			switch event.MergeMessageBody.Message.Data.Type {
			case farcaster.MessageTypeCastAdd.String():
				tasks = append(tasks, &Task{
					Chain:   filter.ChainFarcaster(s.Chain().ID()),
					Message: event.MergeMessageBody.Message,
				})
			case farcaster.MessageTypeReactionAdd.String():
				if event.MergeMessageBody.Message.Data.ReactionBody.Type == farcaster.ReactionTypeRecast.String() {
					tasks = append(tasks, &Task{
						Chain:   filter.ChainFarcaster(s.Chain().ID()),
						Message: event.MergeMessageBody.Message,
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

func (s *source) updateProfileByFid(ctx context.Context, fid *int64) (string, error) {
	username, custodyAddress, ethAddresses, err := s.getProfileByFid(ctx, fid)
	if err != nil {
		return "", fmt.Errorf("get new farcaster profile %d: %w", fid, err)
	}

	if err = s.databaseClient.SaveProfile(ctx, &model.Profile{
		Fid:            *fid,
		Username:       username,
		CustodyAddress: custodyAddress,
		EthAddresses:   ethAddresses,
	}); err != nil {
		return "", err
	}

	return username, nil
}

func (s *source) getProfileByFid(ctx context.Context, fid *int64) (string, string, []string, error) {
	// owner username(handle)
	userDataByFidAndTypeRes, err := s.farcasterClient.GetUserDataByFidAndType(ctx, fid, farcaster.UserDataTypeUsername.String())

	if err != nil {
		return "", "", nil, fmt.Errorf("fetch user data error: %w,%d", err, fid)
	}

	// custody address
	custodyAddresses, err := s.farcasterClient.GetUserNameProofsByFid(ctx, fid)

	if err != nil {
		return "", "", nil, fmt.Errorf("fetch custody address error: %w,%d", err, fid)
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
		return "", "", nil, fmt.Errorf("fetch eth address error: %w,%d", err, fid)
	}

	ethAddresses := lo.Map(verificationsByFidRes.Messages, func(x farcaster.Message, index int) string {
		return common.HexToAddress(x.Data.VerificationAddEthAddressBody.Address).String()
	})

	return userDataByFidAndTypeRes.Data.UserDataBody.Value, common.HexToAddress(custodyAddress).String(), ethAddresses, nil
}

func (s *source) fillMentionsUsernames(ctx context.Context, msg *farcaster.Message) error {
	msg.Data.CastAddBody.MentionsUsernames = make([]string, len(msg.Data.CastAddBody.Mentions))
	for i, fid := range msg.Data.CastAddBody.Mentions {
		fid := int64(fid)
		profile, err := s.databaseClient.LoadProfile(ctx, fid)

		if err != nil {
			zap.L().Warn("fetch farcaster local profile error", zap.Int64("fid", fid))

			continue
		}

		if profile == nil {
			username, err := s.updateProfileByFid(ctx, &fid)
			if err != nil {
				zap.L().Warn("save farcaster local profile error", zap.Int64("fid", fid))

				continue
			}

			profile.Username = username
		}

		msg.Data.CastAddBody.MentionsUsernames[i] = profile.Username
	}

	return nil
}

func NewSource(config *engine.Config, checkpoint *engine.Checkpoint) (engine.Source, error) {
	var state State

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := source{
		config: config,
		state:  state,
	}

	return &instance, nil
}
