package filter

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// Name represents a worker name.
//
//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Name --linecomment --output worker_string.go --json --yaml --sql
type Name int

const (
	Unknown    Name = iota // unknown
	Fallback               // fallback
	Mirror                 // mirror
	Farcaster              // farcaster
	RSS3                   // rss3
	Paragraph              // paragraph
	OpenSea                // opensea
	Uniswap                // uniswap
	Optimism               // optimism
	Aavegotchi             // aavegotchi
	Lens                   // lens
	Looksrare              // looksrare
	Matters                // matters
	Momoka                 // momoka
	Highlight              // highlight
	Aave                   // aave
	IQWiki                 // iqwiki
)

func WorkerHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.Int {
			return data, nil
		}

		return _NameNameToValueMap[data.(string)], nil
	}
}
