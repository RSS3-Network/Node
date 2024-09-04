package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/robfig/cron/v3"
	"github.com/rss3-network/node/config"
	provider "github.com/rss3-network/node/provider/redis"
	"go.uber.org/zap"
)

type CronJob struct {
	crontab *cron.Cron
	mutex   *redsync.Mutex
	timeout time.Duration
}

var KeyPrefix = "cronjob:%s"

func (c *CronJob) AddFunc(ctx context.Context, spec string, cmd func()) error {
	_, err := c.crontab.AddFunc(spec, func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		if err := c.mutex.Lock(); err != nil {
			zap.L().Error("lock error", zap.String("key", c.mutex.Name()), zap.Error(err))

			return
		}

		defer func() {
			if _, err := c.mutex.Unlock(); err != nil {
				zap.L().Error("release lock error", zap.String("key", c.mutex.Name()), zap.Error(err))
			}
		}()

		c.Renewal(ctx)
		cmd()
	})

	return err
}

func (c *CronJob) Renewal(ctx context.Context) {
	go func(ctx context.Context) {
		// Renewal lock every half of timeout.
		ticker := time.NewTicker(c.timeout / 2)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				result, err := c.mutex.ExtendContext(ctx)
				if err != nil {
					zap.L().Error("extend lock error", zap.String("key", c.mutex.Name()), zap.Error(err))

					continue
				}

				if !result {
					zap.L().Error("extend lock failed", zap.String("key", c.mutex.Name()))

					continue
				}
			}
		}
	}(ctx)
}

func (c *CronJob) Start() {
	c.crontab.Start()
}

func (c *CronJob) Stop() {
	c.crontab.Stop()
}

func NewCronJob(configFile *config.File, name string, timeout time.Duration) (*CronJob, error) {
	redisClient, err := provider.NewRedisClient(*configFile.Redis)
	if err != nil {
		return nil, fmt.Errorf("new redis client: %w", err)
	}

	pool := goredis.NewPool(redisClient)
	rs := redsync.New(pool)

	return &CronJob{
		crontab: cron.New(cron.WithLocation(time.UTC), cron.WithSeconds()),
		mutex:   rs.NewMutex(fmt.Sprintf(KeyPrefix, name), redsync.WithExpiry(timeout)),
		timeout: timeout,
	}, nil
}
