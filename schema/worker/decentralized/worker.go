package decentralized

import (
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/protocol-go/schema/tag"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.6 --type=Worker --linecomment --output ../../../docs/schemas/DecentralizedWorker.yaml -t ../../../docs/schemas/tmpl/Decentralized.yaml.tmpl
type Worker int

const (
	Aave       Worker = iota + 1 // aave
	Aavegotchi                   // aavegotchi
	Arbitrum                     // arbitrum
	Base                         // base
	BendDAO                      // benddao
	Core                         // core
	Cow                          // cow
	Crossbell                    // crossbell
	Curve                        // curve
	ENS                          // ens
	Highlight                    // highlight
	IQWiki                       // iqwiki
	KiwiStand                    // kiwistand
	Lens                         // lens
	Lido                         // lido
	Linea                        // linea
	LiNEAR                       // linear
	Looksrare                    // looksrare
	Matters                      // matters
	Mirror                       // mirror
	Momoka                       // momoka
	NearSocial                   // nearsocial
	Nouns                        // nouns
	Oneinch                      // 1inch
	OpenSea                      // opensea
	Optimism                     // optimism
	Paragraph                    // paragraph
	Paraswap                     // paraswap
	Polymarket                   // polymarket
	Rainbow                      // rainbow
	RSS3                         // rss3
	SAVM                         // savm
	Stargate                     // stargate
	Uniswap                      // uniswap
	VSL                          // vsl
	Zerion                       // zerion
)

func (w Worker) Component() string {
	return "decentralized"
}

func (w Worker) Name() string {
	return w.String()
}

var _ echo.BindUnmarshaler = (*Worker)(nil)

func (w *Worker) UnmarshalParam(param string) error {
	worker, err := WorkerString(param)
	if err != nil {
		return err
	}

	*w = worker

	return nil
}

func GetValueByWorkerStr(workerStr string) Worker {
	return _WorkerNameToValueMap[workerStr]
}

// ToTagsMap is a map of worker to tags
var ToTagsMap = map[Worker][]tag.Tag{
	Aave:       {tag.Exchange},
	Aavegotchi: {tag.Metaverse},
	Arbitrum:   {tag.Transaction},
	BendDAO:    {tag.Collectible, tag.Exchange},
	Core:       {tag.Collectible, tag.Transaction},
	Cow:        {tag.Exchange},
	Crossbell:  {tag.Social},
	Curve:      {tag.Exchange, tag.Transaction},
	ENS:        {tag.Social, tag.Collectible},
	Highlight:  {tag.Collectible, tag.Transaction},
	IQWiki:     {tag.Social},
	KiwiStand:  {tag.Collectible, tag.Transaction, tag.Social},
	Lens:       {tag.Social},
	Lido:       {tag.Exchange, tag.Transaction, tag.Collectible},
	Linea:      {tag.Exchange, tag.Transaction},
	LiNEAR:     {tag.Exchange, tag.Transaction},
	Looksrare:  {tag.Collectible},
	Matters:    {tag.Social},
	Mirror:     {tag.Social},
	Momoka:     {tag.Social},
	NearSocial: {tag.Social},
	Nouns:      {tag.Collectible, tag.Governance},
	Oneinch:    {tag.Exchange},
	OpenSea:    {tag.Collectible},
	Optimism:   {tag.Transaction},
	Paragraph:  {tag.Social},
	Paraswap:   {tag.Exchange},
	Polymarket: {tag.Exchange},
	Rainbow:    {tag.Exchange, tag.Transaction},
	RSS3:       {tag.Exchange, tag.Collectible},
	SAVM:       {tag.Transaction},
	Stargate:   {tag.Transaction},
	Uniswap:    {tag.Exchange, tag.Transaction},
	VSL:        {tag.Transaction},
	Zerion:     {tag.Exchange, tag.Transaction},
}
