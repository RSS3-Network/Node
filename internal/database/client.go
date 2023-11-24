package database

import (
	"context"
	"database/sql"
	"embed"

	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	gorm "gorm.io/gorm/schema"
)

//go:embed migration/*.sql
var MigrationFS embed.FS

// Mode is the type of storage mode, which can be either "sharded" or "single".
type Mode string

const (
	ModeSharded Mode = "sharded"
	ModeSingle  Mode = "single"
)

type Table interface {
	gorm.Tabler

	ShardedName() string
}

type Client interface {
	Session
	Transaction

	SaveFeeds(ctx context.Context, feeds []schema.Feed) error
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
