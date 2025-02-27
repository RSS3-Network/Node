package federated

import "github.com/labstack/echo/v4"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Platform --linecomment --output platform_string.go --json --yaml --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.6 --type=Platform --linecomment --output ../../../docs/schemas/FederatedPlatform.yaml -t ../../../docs/schemas/tmpl/Federated.yaml.tmpl
type Platform uint64

const (
	PlatformUnknown  Platform = iota // Unknown
	PlatformMastodon                 // Mastodon
	PlatformBluesky                  // Bluesky
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
	Mastodon: PlatformMastodon,
	Bluesky:  PlatformBluesky,
}
