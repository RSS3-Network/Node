package worker

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
type Worker uint64

const (
	Aave       Worker = iota // aave
	Aavegotchi               // aavegotchi
	Core                     // core
	Crossbell                // crossbell
	Curve                    // curve
	ENS                      // ens
	Highlight                // highlight
	IQWiki                   // iqwiki
	KiwiStand                // kiwistand
	Lens                     // lens
	Lido                     // lido
	Looksrare                // looksrare
	Matters                  // matters
	Mirror                   // mirror
	Momoka                   // momoka
	Oneinch                  // 1inch
	OpenSea                  // opensea
	Optimism                 // optimism
	Paragraph                // paragraph
	RSS3                     // rss3
	SAVM                     // savm
	Stargate                 // stargate
	Uniswap                  // uniswap
	VSL                      // vsl
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
