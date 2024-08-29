package decentralized

import (
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/protocol-go/schema/tag"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
type Worker int

const (
	Aave       Worker = iota + 1 // aave
	Aavegotchi                   // aavegotchi
	Core                         // core
	Crossbell                    // crossbell
	Curve                        // curve
	ENS                          // ens
	Highlight                    // highlight
	IQWiki                       // iqwiki
	KiwiStand                    // kiwistand
	Lens                         // lens
	Lido                         // lido
	Looksrare                    // looksrare
	Matters                      // matters
	Mirror                       // mirror
	Momoka                       // momoka
	Oneinch                      // 1inch
	OpenSea                      // opensea
	Optimism                     // optimism
	Paragraph                    // paragraph
	RSS3                         // rss3
	SAVM                         // savm
	Stargate                     // stargate
	Uniswap                      // uniswap
	VSL                          // vsl
	BendDAO                      // benddao
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
	Core:       {tag.Collectible, tag.Transaction},
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
	BendDAO:    {tag.Collectible, tag.Exchange},
}

var ToIconURLMap = map[Worker]string{
	Aave:       "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/aave.svg",
	Aavegotchi: "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/aavegotchi.svg",
	Core:       "",
	Crossbell:  "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/crossbell.svg",
	Curve:      "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/curve.svg",
	ENS:        "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/ens.svg",
	Highlight:  "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/highlight.svg",
	IQWiki:     "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/iqwiki.svg",
	KiwiStand:  "https://storage.googleapis.com/rss3-icon/kiwistand.png",
	Lens:       "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/lens.svg",
	Lido:       "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/lido.svg",
	Looksrare:  "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/looks-rare.svg",
	Matters:    "https://storage.googleapis.com/rss3-icon/matters.jpg",
	Mirror:     "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/mirror.svg",
	Momoka:     "https://storage.googleapis.com/rss3-icon/momoka.png",
	Oneinch:    "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/1inch.svg",
	OpenSea:    "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/opensea.svg",
	Optimism:   "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/optimism.svg",
	Paragraph:  "https://storage.googleapis.com/rss3-icon/paragraph.png",
	RSS3:       "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/rss3.svg",
	SAVM:       "",
	Stargate:   "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/stargate.svg",
	Uniswap:    "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/uniswap.svg",
	VSL:        "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/rss3-vsl.svg",
	BendDAO:    "https://unpkg.com/@rss3/web3-icons-svg@latest/icons/benddao.svg",
}
