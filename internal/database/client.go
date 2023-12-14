package database

import (
	"context"
	"database/sql"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/mirror/model"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

type Client interface {
	Session
	Transaction
	DatasetMirrorPost

	LoadCheckpoint(ctx context.Context, id string, chain filter.Chain, worker string) (*engine.Checkpoint, error)
	SaveCheckpoint(ctx context.Context, checkpoint *engine.Checkpoint) error

	SaveFeeds(ctx context.Context, feeds []*schema.Feed) error
}

type Session interface {
	Migrate(ctx context.Context) error
	WithTransaction(ctx context.Context, transactionFunction func(ctx context.Context, client Client) error, transactionOptions ...*sql.TxOptions) error
	Begin(ctx context.Context, transactionOptions ...*sql.TxOptions) (Client, error)
}

type Transaction interface {
	Rollback() error
	Commit() error
}

type DatasetMirrorPost interface {
	LoadDatasetMirrorPost(ctx context.Context, originContentDigest string) (*model.DatasetMirrorPost, error)
	SaveDatasetMirrorPost(ctx context.Context, post *model.DatasetMirrorPost) error
}

var _ goose.Logger = (*SugaredLogger)(nil)

type SugaredLogger struct {
	Logger *zap.SugaredLogger
}

func (s SugaredLogger) Fatalf(format string, v ...interface{}) {
	s.Logger.Fatalf(format, v...)
}

func (s SugaredLogger) Printf(format string, v ...interface{}) {
	s.Logger.Infof(format, v...)
}
