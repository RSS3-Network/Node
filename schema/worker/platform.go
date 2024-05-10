package worker

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Platform --linecomment --output platform_string.go --json --yaml --sql
type Platform uint64

const (
	Platform1Inch      Platform = iota // 1inch
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
