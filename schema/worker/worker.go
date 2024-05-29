package worker

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/rss3-network/protocol-go/schema/tag"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
type Worker int

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
	RSSHub                   // rsshub
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

		if t.Kind() != reflect.Int {
			return data, nil
		}

		return _WorkerNameToValueMap[data.(string)], nil
	}
}

// ToTagsMap is a map of worker to tags
var ToTagsMap = map[Worker][]tag.Tag{
	Aave:       {tag.Exchange},
	Aavegotchi: {tag.Metaverse},
	Crossbell:  {tag.Social},
	Curve:      {tag.Exchange, tag.Transaction},
	ENS:        {tag.Social, tag.Collectible},
	Highlight:  {tag.Collectible, tag.Transaction},
	IQWiki:     {tag.Social},
	KiwiStand:  {tag.Collectible, tag.Transaction, tag.Social},
	Lens:       {tag.Social},
	Lido:       {tag.Exchange, tag.Transaction, tag.Collectible},
	Looksrare:  {tag.Collectible},
	Matters:    {tag.Social},
	Mirror:     {tag.Social},
	Momoka:     {tag.Social},
	Oneinch:    {tag.Exchange},
	OpenSea:    {tag.Collectible},
	Optimism:   {tag.Transaction},
	Paragraph:  {tag.Social},
	RSS3:       {tag.Exchange, tag.Collectible},
	SAVM:       {tag.Transaction},
	Stargate:   {tag.Transaction},
	Uniswap:    {tag.Exchange, tag.Transaction},
	VSL:        {tag.Transaction},
}
