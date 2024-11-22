package decentralized

import "github.com/labstack/echo/v4"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Platform --linecomment --output platform_string.go --json --yaml --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.6 --type=Platform --linecomment --output ../../../docs/schemas/PlatformDecentralized.yaml -t ../../../docs/schemas/tmpl/Decentralized.yaml.tmpl
type Platform uint64

const (
	PlatformUnknown    Platform = iota // Unknown
	Platform1Inch                      // 1inch
	PlatformAAVE                       // AAVE
	PlatformAavegotchi                 // Aavegotchi
	PlatformArbitrum                   // Arbitrum
	PlatformBase                       // Base
	PlatformBendDAO                    // BendDAO
	PlatformCow                        // Cow
	PlatformCrossbell                  // Crossbell
	PlatformCurve                      // Curve
	PlatformENS                        // ENS
	PlatformFarcaster                  // Farcaster
	PlatformHighlight                  // Highlight
	PlatformIQWiki                     // IQWiki
	PlatformKiwiStand                  // KiwiStand
	PlatformLens                       // Lens
	PlatformLido                       // Lido
	PlatformLinea                      // Linea
	PlatformLiNEAR                     // LiNEAR
	PlatformLooksRare                  // LooksRare
	PlatformMatters                    // Matters
	PlatformMirror                     // Mirror
	PlatformNearSocial                 // NearSocial
	PlatformNouns                      // Nouns
	PlatformOpenSea                    // OpenSea
	PlatformOptimism                   // Optimism
	PlatformParagraph                  // Paragraph
	PlatformParaswap                   // Paraswap
	PlatformPolymarket                 // Polymarket
	PlatformRSS3                       // RSS3
	PlatformRainbow                    // Rainbow
	PlatformSAVM                       // SAVM
	PlatformStargate                   // Stargate
	PlatformUniswap                    // Uniswap
	PlatformVSL                        // VSL
	PlatformZerion                     // Zerion
)

var _ echo.BindUnmarshaler = (*Platform)(nil)

func (p *Platform) UnmarshalParam(param string) error {
	platform, err := PlatformString(param)
	if err != nil {
		return err
	}

	*p = platform

	return nil
}

// ToPlatformMap is a map of worker to platform
var ToPlatformMap = map[Worker]Platform{
	Aave:       PlatformAAVE,
	Aavegotchi: PlatformAavegotchi,
	Arbitrum:   PlatformArbitrum,
	Base:       PlatformBase,
	BendDAO:    PlatformBendDAO,
	Cow:        PlatformCow,
	Crossbell:  PlatformCrossbell,
	Curve:      PlatformCurve,
	ENS:        PlatformENS,
	Highlight:  PlatformHighlight,
	IQWiki:     PlatformIQWiki,
	KiwiStand:  PlatformKiwiStand,
	Lens:       PlatformLens,
	Lido:       PlatformLido,
	Linea:      PlatformLinea,
	LiNEAR:     PlatformLiNEAR,
	Looksrare:  PlatformLooksRare,
	Matters:    PlatformMatters,
	Mirror:     PlatformMirror,
	Momoka:     PlatformLens,
	NearSocial: PlatformNearSocial,
	Nouns:      PlatformNouns,
	Oneinch:    Platform1Inch,
	OpenSea:    PlatformOpenSea,
	Optimism:   PlatformOptimism,
	Paragraph:  PlatformParagraph,
	Paraswap:   PlatformParaswap,
	Polymarket: PlatformPolymarket,
	Rainbow:    PlatformRainbow,
	RSS3:       PlatformRSS3,
	SAVM:       PlatformSAVM,
	Stargate:   PlatformStargate,
	Uniswap:    PlatformUniswap,
	VSL:        PlatformVSL,
	Zerion:     PlatformZerion,
}
