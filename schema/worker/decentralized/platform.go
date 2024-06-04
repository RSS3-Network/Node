package decentralized

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Platform --linecomment --output platform_string.go --json --yaml --sql
type Platform uint64

const (
	PlatformUnknown    Platform = iota // Unknown
	Platform1Inch                      // 1inch
	PlatformAAVE                       // AAVE
	PlatformAavegotchi                 // Aavegotchi
	PlatformCrossbell                  // Crossbell
	PlatformCurve                      // Curve
	PlatformENS                        // ENS
	PlatformFarcaster                  // Farcaster
	PlatformHighlight                  // Highlight
	PlatformIQWiki                     // IQWiki
	PlatformKiwiStand                  // KiwiStand
	PlatformLens                       // Lens
	PlatformLido                       // Lido
	PlatformLooksRare                  // LooksRare
	PlatformMatters                    // Matters
	PlatformMirror                     // Mirror
	PlatformOpenSea                    // OpenSea
	PlatformOptimism                   // Optimism
	PlatformParagraph                  // Paragraph
	PlatformRSS3                       // RSS3
	PlatformSAVM                       // SAVM
	PlatformStargate                   // Stargate
	PlatformUniswap                    // Uniswap
	PlatformVSL                        // VSL
)

// ToPlatformMap is a map of worker to platform
var ToPlatformMap = map[Worker]Platform{
	Aave:       PlatformAAVE,
	Aavegotchi: PlatformAavegotchi,
	Crossbell:  PlatformCrossbell,
	Curve:      PlatformCurve,
	ENS:        PlatformENS,
	Highlight:  PlatformHighlight,
	IQWiki:     PlatformIQWiki,
	KiwiStand:  PlatformKiwiStand,
	Lens:       PlatformLens,
	Lido:       PlatformLido,
	Looksrare:  PlatformLooksRare,
	Matters:    PlatformMatters,
	Mirror:     PlatformMirror,
	Momoka:     PlatformLens,
	Oneinch:    Platform1Inch,
	OpenSea:    PlatformOpenSea,
	Optimism:   PlatformOptimism,
	Paragraph:  PlatformParagraph,
	RSS3:       PlatformRSS3,
	SAVM:       PlatformSAVM,
	Stargate:   PlatformStargate,
	Uniswap:    PlatformUniswap,
	VSL:        PlatformVSL,
}
