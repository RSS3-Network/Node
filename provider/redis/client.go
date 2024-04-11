package redis

import (
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
)

func NewClient(option config.Redis) (rueidis.Client, error) {
	clientOption := rueidis.ClientOption{
		InitAddress:  option.Endpoints,
		Username:     option.Username,
		Password:     option.Password,
		DisableCache: option.DisableCache,
	}

	return rueidis.NewClient(clientOption)
}
