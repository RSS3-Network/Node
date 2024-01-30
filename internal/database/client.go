package database

import (
	"context"
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pressly/goose/v3"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/engine"
	mirror_model "github.com/rss3-network/node/internal/engine/worker/contract/mirror/model"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"go.uber.org/zap"
)

type Client interface {
	Session
	Transaction
	DatasetMirrorPost
	DatasetFarcasterProfile
	DatasetENSNamehash

	LoadCheckpoint(ctx context.Context, id string, network filter.Network, worker string) (*engine.Checkpoint, error)
	LoadCheckpoints(ctx context.Context, id string, network filter.Network, worker string) ([]*engine.Checkpoint, error)
	SaveCheckpoint(ctx context.Context, checkpoint *engine.Checkpoint) error

	SaveFeeds(ctx context.Context, feeds []*schema.Feed) error
	FindFeed(ctx context.Context, query model.FeedQuery) (*schema.Feed, *int, error)
	FindFeeds(ctx context.Context, query model.FeedsQuery) ([]*schema.Feed, error)
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
	LoadDatasetMirrorPost(ctx context.Context, originContentDigest string) (*mirror_model.DatasetMirrorPost, error)
	SaveDatasetMirrorPost(ctx context.Context, post *mirror_model.DatasetMirrorPost) error
}

type DatasetFarcasterProfile interface {
	LoadDatasetFarcasterProfile(ctx context.Context, fid int64) (*model.Profile, error)
	SaveDatasetFarcasterProfile(ctx context.Context, profile *model.Profile) error
}

type DatasetENSNamehash interface {
	LoadDatasetENSNamehash(ctx context.Context, hash common.Hash) (*model.ENSNamehash, error)
	SaveDatasetENSNamehash(ctx context.Context, namehash *model.ENSNamehash) error
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
