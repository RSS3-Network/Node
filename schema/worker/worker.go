package worker

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/rss"
)

type Worker interface {
	Component() string
	Name() string
}

func HookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String || t.Kind() != reflect.TypeOf((*Worker)(nil)).Elem().Kind() {
			return data, nil
		}

		// TODO: Implement the logic to determine the worker type
		if value := rss.GetValueByWorkerStr(data.(string)); value != 0 {
			return value, nil
		}

		return decentralized.GetValueByWorkerStr(data.(string)), nil
	}
}
