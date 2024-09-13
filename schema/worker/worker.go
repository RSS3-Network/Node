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
		// Only process if the target type is Worker and the source type is string
		if f.Kind() == reflect.String && t.Kind() == reflect.TypeOf((*Worker)(nil)).Elem().Kind() {
			workerStr := data.(string)

			// TODO: Implement the logic to determine the worker type
			if value := rss.GetValueByWorkerStr(workerStr); value != 0 {
				return value, nil
			} else if value := decentralized.GetValueByWorkerStr(workerStr); value != 0 {
				return value, nil
			}
		}
		// For all other types, return data
		return data, nil
	}
}
