package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
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
	state           State
	pendingState    State
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

		s.state.CastsFid = fidsResponse.Fids[0]
	}

	for s.state.CastsFid > 1 {
		err := s.pollCastsByFid(ctx, lo.ToPtr(int64(s.state.CastsFid)), "", tasksChan)

		if err != nil {
			return err
		}

		// TODO update checkpoint

		s.state = s.pendingState
		s.pendingState.CastsFid--
	}

	return nil
}

func (s *source) pollCastsByFid(ctx context.Context, fid *int64, pageToken string, tasksChan chan<- []engine.Task) error {
	for {
		castsByFidResponse, err := s.farcasterClient.GetCastsByFid(ctx, fid, true, nil, pageToken)

		if err != nil {
			return err
		}

		tasks, err := s.buildFarcasterMessageTasks(castsByFidResponse.Messages)
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

		s.state.ReactionsFid = fidsResponse.Fids[0]
	}

	for s.state.ReactionsFid > 1 {
		err := s.pollReactionsByFid(ctx, lo.ToPtr(int64(s.state.ReactionsFid)), "", tasksChan)

		if err != nil {
			return err
		}

		// TODO update checkpoint

		s.state = s.pendingState
		s.pendingState.ReactionsFid--
	}

	return nil
}

func (s *source) pollReactionsByFid(ctx context.Context, fid *int64, pageToken string, tasksChan chan<- []engine.Task) error {
	for {
		reactionsByFidResponse, err := s.farcasterClient.GetReactionsByFid(ctx, fid, true, nil, pageToken, farcaster.ReactionTypeRecast.String())

		if err != nil {
			return err
		}

		tasks, err := s.buildFarcasterMessageTasks(reactionsByFidResponse.Messages)
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

func (s *source) buildFarcasterMessageTasks(messages []farcaster.Message) ([]*Task, error) {
	tasks := make([]*Task, 0)

	for _, msg := range messages {
		if msg.Data.Type == farcaster.MessageTypeCastAdd.String() ||
			(msg.Data.Type == farcaster.MessageTypeReactionAdd.String() && msg.Data.ReactionBody.Type == farcaster.ReactionTypeRecast.String()) {
			tasks = append(tasks, &Task{
				Chain:   filter.ChainFarcaster(s.Chain().ID()),
				Message: msg,
			})
		}
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

		// TODO update checkpoint

		s.state = s.pendingState
		s.pendingState.EventID = eventsResponse.NextPageEventID

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

				// TODO update table profile

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

func NewSource(config *engine.Config, checkpoint *engine.Checkpoint) (engine.Source, error) {
	var state State

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := source{
		config:       config,
		state:        state,
		pendingState: state, // Default pending state is equal to current state.
	}

	return &instance, nil
}
