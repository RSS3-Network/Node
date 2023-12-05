package farcaster

import "time"

type State struct {
	EventID        uint64    `json:"event_id"`
	EventTimestamp time.Time `json:"event_timestamp"`
	CastsFid       uint64    `json:"casts_fid"`
	ReactionsFid   uint64    `json:"reactions_fid"`
}
