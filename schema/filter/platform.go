package filter

import (
	"strings"

	"github.com/labstack/echo/v4"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Platform --linecomment --output platform_string.go --json --sql
type Platform int

const (
	PlatformRSS3       Platform = iota + 1 // RSS3
	PlatformMirror                         // Mirror
	PlatformFarcaster                      // Farcaster
	PlatformParagraph                      // Paragraph
	PlatformOpenSea                        // OpenSea
	PlatformUniswap                        // Uniswap
	PlatformOptimism                       // Optimism
	PlatformAavegotchi                     // Aavegotchi
	PlatformLens                           // Lens
	PlatformLooksRare                      // LooksRare
	PlatformCrossbell                      // Crossbell
)

func (p Platform) ID() string {
	return strings.ReplaceAll(strings.ToLower(p.String()), "\x20", "_")
}

var _ echo.BindUnmarshaler = (*Platform)(nil)

func (p *Platform) UnmarshalParam(param string) error {
	platform, err := PlatformString(param)
	if err != nil {
		return err
	}

	*p = platform

	return nil
}
