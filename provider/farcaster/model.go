package farcaster

import "github.com/rss3-network/node/internal/database/model"

const FarcasterEpoch = 1609459200 // January 1, 2021 UTC https://github.com/farcasterxyz/hub-monorepo/blob/77ff79ed804104956eb153247c22c00099c7b122/packages/core/src/time.ts#L4

type farcasterQuery struct {
	Fid          *int64 `form:"fid,omitempty"`
	TargetFid    *int64 `form:"target_fid,omitempty"`
	Hash         string `form:"hash,omitempty"`
	TargetHash   string `form:"target_hash,omitempty"`
	FromEventID  string `form:"from_event_id,omitempty"`
	Reverse      bool   `form:"reverse,omitempty"`
	PageSize     *int   `form:"pageSize,omitempty"`
	PageToken    string `form:"pageToken,omitempty"`
	UserDataType string `form:"user_data_type,omitempty"`
	ReactionType string `form:"reaction_type,omitempty"`
}

type MessageResponse struct {
	Messages      []Message `json:"messages"`
	NextPageToken string    `json:"nextPageToken"`
}

type ProofResponse struct {
	Proofs []UserNameProof `json:"proofs"`
}

type EventResponse struct {
	NextPageEventID uint64     `json:"nextPageEventId"`
	Events          []HubEvent `json:"events"`
}

type FidResponse struct {
	Fids          []uint64 `json:"fids"`
	NextPageToken string   `json:"nextPageToken"`
}

type Message struct {
	Data            MessageData `json:"data"`
	Hash            string      `json:"hash"`
	HashScheme      string      `json:"hashScheme"`
	Signature       string      `json:"signature"`
	SignatureScheme string      `json:"signatureScheme"`
	Signer          string      `json:"signer"`
}

type MessageData struct {
	Type                          string                         `json:"type"`
	Fid                           uint64                         `json:"fid"`
	Profile                       *model.Profile                 `json:"profile,omitempty"`
	Timestamp                     uint32                         `json:"timestamp"`
	Network                       string                         `json:"network"`
	CastAddBody                   *CastAddBody                   `json:"castAddBody,omitempty"`
	CastRemoveBody                *CastRemoveBody                `json:"castRemoveBody,omitempty"`
	UserDataBody                  *UserDataBody                  `json:"userDataBody,omitempty"`
	ReactionBody                  *ReactionBody                  `json:"reactionBody,omitempty"`
	LinkBody                      *LinkBody                      `json:"linkBody,omitempty"`
	VerificationAddEthAddressBody *VerificationAddEthAddressBody `json:"verificationAddEthAddressBody,omitempty"`
	VerificationRemoveBody        *VerificationRemoveBody        `json:"verificationRemoveBody,omitempty"`
	UserNameProof                 *UserNameProof                 `json:"userNameProof,omitempty"`
}

type CastAddBody struct {
	EmbedsDeprecated  []string `json:"embedsDeprecated"`
	Mentions          []uint64 `json:"mentions"`
	MentionsUsernames []string `json:"mentionsUsernames"`
	ParentCastID      *CastID  `json:"parentCastId,omitempty"`
	ParentCast        *Message `json:"parentCast,omitempty"`
	ParentURL         string   `json:"parentUrl,omitempty"`
	Text              string   `json:"text"`
	MentionsPositions []int32  `json:"mentionsPositions"`
	Embeds            []Embed  `json:"embeds"`
}

type CastID struct {
	Fid  uint64 `json:"fid"`
	Hash string `json:"hash"`
}

type Embed struct {
	URL string `json:"url"`
}

type CastRemoveBody struct {
	TargetHash string `json:"targeHash"`
}

type UserDataBody struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ReactionBody struct {
	Type         string   `json:"type"`
	TargetCastID CastID   `json:"targetCastId"`
	TargetCast   *Message `json:"targetCast"`
	TargetURL    string   `json:"targetUrl"`
}

type LinkBody struct {
	Type             string `json:"type"`
	DisplayTimestamp uint32 `json:"displayTimestamp"`
	TargetFid        uint64 `json:"targetFid"`
}

type VerificationAddEthAddressBody struct {
	Address      string `json:"address"`
	EthSignature string `json:"ethSignature"`
	BlockHash    string `json:"blockHash"`
}

type VerificationRemoveBody struct {
	Address string `json:"address"`
}

type UserNameProof struct {
	Timestamp uint32 `json:"timestamp"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	Signature string `json:"signature"`
	Fid       uint64 `json:"fid"`
	Type      string `json:"type"`
}

type HubEvent struct {
	Type                   string                  `json:"type"`
	ID                     uint64                  `json:"id"`
	MergeMessageBody       *MergeMessageBody       `json:"mergeMessageBody,omitempty"`
	PruneMessageBody       *PruneMessageBody       `json:"pruneMessageBody,omitempty"`
	RevokeMessageBody      *RevokeMessageBody      `json:"revokeMessageBody,omitempty"`
	MergeUserNameProofBody *MergeUserNameProofBody `json:"mergeUserNameProofBody,omitempty"`
	MergeOnChainEventBody  *MergeOnChainEventBody  `json:"mergeOnChainEventBody,omitempty"`
}

type MergeMessageBody struct {
	Message         Message   `json:"message"`
	DeletedMessages []Message `json:"deletedMessages"`
}

type PruneMessageBody struct {
	Message Message `json:"message"`
}

type RevokeMessageBody struct {
	Message Message `json:"message"`
}

type MergeUserNameProofBody struct {
	UserNameProof               UserNameProof `json:"usernameProof"`
	DeletedUserNameProof        UserNameProof `json:"deletedUsernameProof"`
	UsernameProofMessage        Message       `json:"usernameProofMessage"`
	DeletedUsernameProofMessage Message       `json:"deletedUsernameProofMessage"`
}

type MergeOnChainEventBody struct{}

// Profile redis profile
type Profile struct {
	Username       string
	CustodyAddress string
	EthAddresses   []string
}
