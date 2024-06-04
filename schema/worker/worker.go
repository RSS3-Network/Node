package worker

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/rss"
)

type Worker interface {
	Name() string
}

func HookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.TypeOf((*Worker)(nil)).Elem().Kind() {
			return data, nil
		}

		// TODO: Implement the logic to determine the worker type
		if data.(string) == rss.Rsshub.String() {
			return rss.GetValueByWorkerStr(data.(string)), nil
		}

		return decentralized.GetValueByWorkerStr(data.(string)), nil
	}
}
