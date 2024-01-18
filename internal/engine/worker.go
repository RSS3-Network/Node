package engine

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/schema"
)

// Name represents a worker name.
//
//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Name --linecomment --output worker_string.go --json --yaml --sql
type Name int

const (
	Unknown    Name = iota // unknown
	Fallback               // fallback
	Mirror                 // mirror
	Farcaster              // farcaster
	RSS3                   // rss3
	Paragraph              // paragraph
	OpenSea                // opensea
	Uniswap                // uniswap
	Optimism               // optimism
	Aavegotchi             // aavegotchi
	Lens                   // lens
	Matters                // matters
)

type Worker interface {
	Name() string
	Filter() SourceFilter
	Match(ctx context.Context, task Task) (bool, error)
	Transform(ctx context.Context, task Task) (*schema.Feed, error)
}
