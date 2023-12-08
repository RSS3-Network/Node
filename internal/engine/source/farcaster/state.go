package farcaster

type State struct {
	EventID           uint64 `json:"event_id"`
	CastsFid          uint64 `json:"casts_fid"`
	CastsBackfill     bool   `json:"casts_backfill"`
	ReactionsFid      uint64 `json:"reactions_fid"`
	ReactionsBackfill bool   `json:"reactions_backfill"`
}
