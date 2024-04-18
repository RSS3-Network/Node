package worker

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
type Worker uint64

const (
	Fallback   Worker = iota // fallback
	Mirror                   // mirror
	Farcaster                // farcaster
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
	Oneinch                  // oneinch
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

const (
	platform1Inch      string = "1inch"
	platformAAVE       string = "AAVE"
	platformAavegotchi string = "Aavegotchi"
	platformCrossbell  string = "Crossbell"
	platformCurve      string = "Curve"
	platformENS        string = "ENS"
	platformFarcaster  string = "Farcaster"
	platformHighlight  string = "Highlight"
	platformIQWiki     string = "IQWiki"
	platformKiwiStand  string = "KiwiStand"
	platformLens       string = "Lens"
	platformLido       string = "Lido"
	platformLooksRare  string = "LooksRare"
	platformMatters    string = "Matters"
	platformMirror     string = "Mirror"
	platformOpenSea    string = "OpenSea"
	platformOptimism   string = "Optimism"
	platformParagraph  string = "Paragraph"
	platformRSS3       string = "RSS3"
	platformSAVM       string = "SAVM"
	platformStargate   string = "Stargate"
	platformUniswap    string = "Uniswap"
	platformVSL        string = "VSL"
)

// Platform returns the display name of the worker.
func (w Worker) Platform() string {
	switch w {
	case Mirror:
		return platformMirror
	case Farcaster:
		return platformFarcaster
	case RSS3:
		return platformRSS3
	case Paragraph:
		return platformParagraph
	case OpenSea:
		return platformOpenSea
	case Uniswap:
		return platformUniswap
	case Optimism:
		return platformOptimism
	case Aavegotchi:
		return platformAavegotchi
	case Lens:
		return platformLens
	case Looksrare:
		return platformLooksRare
	case Matters:
		return platformMatters
	case Momoka:
		return platformLens
	case Highlight:
		return platformHighlight
	case Aave:
		return platformAAVE
	case IQWiki:
		return platformIQWiki
	case Lido:
		return platformLido
	case Crossbell:
		return platformCrossbell
	case ENS:
		return platformENS
	case Oneinch:
		return platform1Inch
	case KiwiStand:
		return platformKiwiStand
	case SAVM:
		return platformSAVM
	case VSL:
		return platformVSL
	case Stargate:
		return platformStargate
	case Curve:
		return platformCurve
	default:
		return ""
	}
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