package atproto

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/events"
	"github.com/bluesky-social/indigo/events/schedulers/sequential"
	"github.com/gorilla/websocket"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	at "github.com/rss3-network/node/provider/atproto"
	"github.com/rss3-network/node/provider/atproto/bluesky"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

var _ engine.DataSource = (*dataSource)(nil)

type dataSource struct {
	config         *config.Module
	filter         *Filter
	state          State
	option         *Option
	client         *bluesky.Client
	databaseClient database.Client
}

func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

// Start initializes the data source and starts polling for events and historical repos.
func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("initialize dataSource failed: %w", err)

		return
	}

	// Get latest events
	go func() {
		if err := s.pollSubscribeRepos(ctx, tasksChan); err != nil {
			zap.L().Error("poll subscribe repos failed", zap.Error(err))

			errorChan <- fmt.Errorf("poll subscribe repos failed: %w", err)
		}
	}()

	// Get historical repos
	go func() {
		if err := s.pollListRepos(ctx, tasksChan); err != nil {
			zap.L().Error("poll list repos failed", zap.Error(err))

			errorChan <- fmt.Errorf("poll list repos failed: %w", err)
		}
	}()
}

// pollSubscribeRepos subscribes to repository events and processes them.
func (s *dataSource) pollSubscribeRepos(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	uri := bluesky.BskySubscribeURI

	if s.state.SubscribeCursor != "" {
		uri = fmt.Sprintf("%s?cursor=%s", uri, s.state.SubscribeCursor)
	}

	conn, _, err := websocket.DefaultDialer.Dial(uri, nil)
	if err != nil {
		zap.L().Error("dial websocket failed", zap.Error(err))

		return fmt.Errorf("dial websocket failed: %w", err)
	}

	defer func() {
		zap.L().Info("close websocket connection")

		_ = conn.Close()
	}()

	rsc := &events.RepoStreamCallbacks{
		RepoCommit: func(evt *atproto.SyncSubscribeRepos_Commit) error {
			messages := make([]*at.Message, 0)

			for _, op := range evt.Ops {
				if op.Cid == nil {
					continue
				}

				// Parse the operation
				message, err := s.client.GetRepoRecord(ctx, evt.Repo, op.Path)
				if err != nil {
					zap.L().Error("get subscribe repo record failed", zap.Error(err), zap.String("repo", evt.Repo), zap.Any("op", op))

					return fmt.Errorf("get subscribe repo record failed: %w", err)
				}

				if message != nil {
					messages = append(messages, message)
				}
			}

			if len(messages) > 0 {
				tasksChan <- s.buildTasks(ctx, messages)
			}

			return nil
		},
	}

	sched := sequential.NewScheduler("myfirehose", rsc.EventHandler)

	if err = events.HandleRepoStream(context.Background(), conn, sched); err != nil {
		zap.L().Error("handle repo stream failed", zap.Error(err))

		return fmt.Errorf("handle repo stream failed: %w", err)
	}

	return nil
}

// pollListRepos polls the list of repositories and processes them.
func (s *dataSource) pollListRepos(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	for {
		var cursor string

		if lo.IsNotEmpty(s.state.ListReposCursor) {
			cursor = s.state.ListReposCursor
		}

		repos, next, err := s.client.SyncListRepos(ctx, cursor, bluesky.SyncListReposLimit)
		if err != nil {
			zap.L().Error("sync list repos failed", zap.Error(err))

			return fmt.Errorf("sync list repos failed: %w", err)
		}

		if len(repos) > 0 {
			tasksChan <- s.buildTasks(ctx, repos)
		}

		s.state.ListReposCursor = lo.FromPtr(next)
	}
}

// buildTasks constructs tasks from the given messages.
func (s *dataSource) buildTasks(_ context.Context, messages []*at.Message) *engine.Tasks {
	var tasks engine.Tasks

	if len(messages) == 0 {
		return &tasks
	}

	for _, message := range messages {
		task := &Task{
			Network: s.config.Network,
			Message: lo.FromPtr(message),
		}

		tasks.Tasks = append(tasks.Tasks, task)
	}

	zap.L().Debug("build event tasks", zap.Any("tasks", tasks))

	return &tasks
}

// initialize creates a new Bluesky client and assigns it to the data source.
func (s *dataSource) initialize() error {
	client, err := bluesky.NewClient(context.Background(), s.filter.Type, s.option.Username, s.option.Password)
	if err != nil {
		zap.L().Error("create bluesky client failed", zap.Error(err))

		return fmt.Errorf("create bluesky client failed: %w", err)
	}

	s.client = client

	return nil
}

// NewSource creates a new data source instance with the given configuration, filter, checkpoint, and database client.
func NewSource(config *config.Module, sourceFilter engine.DataSourceFilter, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.DataSource, error) {
	var (
		state State
		err   error
	)

	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, fmt.Errorf("unmarshal checkpoint state failed: %w", err)
		}
	}

	instance := dataSource{
		config:         config,
		state:          state,
		filter:         new(Filter),
		databaseClient: databaseClient,
	}

	if instance.option, err = NewOption(config.Parameters); err != nil {
		return nil, fmt.Errorf("create option failed: %w", err)
	}

	if sourceFilter != nil {
		var ok bool

		if instance.filter, ok = sourceFilter.(*Filter); !ok {
			return nil, fmt.Errorf("invalid filter type")
		}
	}

	return &instance, nil
}
