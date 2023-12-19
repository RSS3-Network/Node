package filter

import "strings"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Platform --linecomment --output platform_string.go --json --sql
type Platform int

func (p Platform) ID() string {
	return strings.ReplaceAll(strings.ToLower(p.String()), "\x20", "_")
}

const (
	PlatformRSS3      Platform = iota + 1 // RSS3
	PlatformMirror                        // Mirror
	PlatformFarcaster                     // Farcaster
)
