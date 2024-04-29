package worker

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
type Worker uint64

const (
	Core       Worker = iota // core
	Mirror                   // mirror
	RSS3                     // rss3
	Paragraph                // paragraph
	OpenSea                  // opensea
	Uniswap                  // uniswapz
	Optimism                 // optimism
	Aavegotchi               // aavegotchi
	Lens                     // lens
	Looksrare                // looksrare
	Matters                  // matters
	Momoka                   // momoka
	Highlight                // highlight
	Aave                     // aave
	IQWiki                   // iqwiki
	Lido                     // lido
	Crossbell                // crossbell
	ENS                      // ens
	Oneinch                  // 1inch
	KiwiStand                // kiwistand
	SAVM                     // savm
	VSL                      // vsl
	Stargate                 // stargate
	Curve                    // curve
)

var _ echo.BindUnmarshaler = (*Worker)(nil)

func (w *Worker) UnmarshalParam(param string) error {
	worker, err := WorkerString(param)
	if err != nil {
		return err
	}

	*w = worker

	return nil
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

		if t.Kind() != reflect.Uint64 {
			return data, nil
		}

		return _WorkerNameToValueMap[data.(string)], nil
	}
}
