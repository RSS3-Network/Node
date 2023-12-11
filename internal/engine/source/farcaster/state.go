package farcaster

type State struct {
	EventID           uint64 `json:"event_id"`           // Event ID that has been processed
	CastsFid          uint64 `json:"casts_fid"`          // Casts ID that has been processed in backfill casts
	CastsBackfill     bool   `json:"casts_backfill"`     // Casts backfill flag
	ReactionsFid      uint64 `json:"reactions_fid"`      // Reactions ID that has been processed in backfill reactions
	ReactionsBackfill bool   `json:"reactions_backfill"` // Reactions backfill flag
}
