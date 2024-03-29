// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package snapshot

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// ProposalProposalsProposal includes the requested fields of the GraphQL type Proposal.
type ProposalProposalsProposal struct {
	Id      string                         `json:"id"`
	Link    string                         `json:"link"`
	Title   string                         `json:"title"`
	Body    string                         `json:"body"`
	Choices []string                       `json:"choices"`
	Start   int                            `json:"start"`
	End     int                            `json:"end"`
	State   string                         `json:"state"`
	Author  string                         `json:"author"`
	Created int                            `json:"created"`
	Space   ProposalProposalsProposalSpace `json:"space"`
}

// GetId returns ProposalProposalsProposal.Id, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetId() string { return v.Id }

// GetLink returns ProposalProposalsProposal.Link, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetLink() string { return v.Link }

// GetTitle returns ProposalProposalsProposal.Title, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetTitle() string { return v.Title }

// GetBody returns ProposalProposalsProposal.Body, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetBody() string { return v.Body }

// GetChoices returns ProposalProposalsProposal.Choices, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetChoices() []string { return v.Choices }

// GetStart returns ProposalProposalsProposal.Start, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetStart() int { return v.Start }

// GetEnd returns ProposalProposalsProposal.End, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetEnd() int { return v.End }

// GetState returns ProposalProposalsProposal.State, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetState() string { return v.State }

// GetAuthor returns ProposalProposalsProposal.Author, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetAuthor() string { return v.Author }

// GetCreated returns ProposalProposalsProposal.Created, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetCreated() int { return v.Created }

// GetSpace returns ProposalProposalsProposal.Space, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposal) GetSpace() ProposalProposalsProposalSpace { return v.Space }

// ProposalProposalsProposalSpace includes the requested fields of the GraphQL type Space.
type ProposalProposalsProposalSpace struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

// GetId returns ProposalProposalsProposalSpace.Id, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposalSpace) GetId() string { return v.Id }

// GetName returns ProposalProposalsProposalSpace.Name, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposalSpace) GetName() string { return v.Name }

// GetAbout returns ProposalProposalsProposalSpace.About, and is useful for accessing the field via an interface.
func (v *ProposalProposalsProposalSpace) GetAbout() string { return v.About }

// ProposalResponse is returned by Proposal on success.
type ProposalResponse struct {
	Proposals []ProposalProposalsProposal `json:"proposals"`
}

// GetProposals returns ProposalResponse.Proposals, and is useful for accessing the field via an interface.
func (v *ProposalResponse) GetProposals() []ProposalProposalsProposal { return v.Proposals }

// ProposalsProposalsProposal includes the requested fields of the GraphQL type Proposal.
type ProposalsProposalsProposal struct {
	Id      string                          `json:"id"`
	Link    string                          `json:"link"`
	Title   string                          `json:"title"`
	Body    string                          `json:"body"`
	Choices []string                        `json:"choices"`
	Start   int                             `json:"start"`
	End     int                             `json:"end"`
	State   string                          `json:"state"`
	Author  string                          `json:"author"`
	Created int                             `json:"created"`
	Space   ProposalsProposalsProposalSpace `json:"space"`
}

// GetId returns ProposalsProposalsProposal.Id, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetId() string { return v.Id }

// GetLink returns ProposalsProposalsProposal.Link, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetLink() string { return v.Link }

// GetTitle returns ProposalsProposalsProposal.Title, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetTitle() string { return v.Title }

// GetBody returns ProposalsProposalsProposal.Body, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetBody() string { return v.Body }

// GetChoices returns ProposalsProposalsProposal.Choices, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetChoices() []string { return v.Choices }

// GetStart returns ProposalsProposalsProposal.Start, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetStart() int { return v.Start }

// GetEnd returns ProposalsProposalsProposal.End, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetEnd() int { return v.End }

// GetState returns ProposalsProposalsProposal.State, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetState() string { return v.State }

// GetAuthor returns ProposalsProposalsProposal.Author, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetAuthor() string { return v.Author }

// GetCreated returns ProposalsProposalsProposal.Created, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetCreated() int { return v.Created }

// GetSpace returns ProposalsProposalsProposal.Space, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposal) GetSpace() ProposalsProposalsProposalSpace { return v.Space }

// ProposalsProposalsProposalSpace includes the requested fields of the GraphQL type Space.
type ProposalsProposalsProposalSpace struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

// GetId returns ProposalsProposalsProposalSpace.Id, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposalSpace) GetId() string { return v.Id }

// GetName returns ProposalsProposalsProposalSpace.Name, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposalSpace) GetName() string { return v.Name }

// GetAbout returns ProposalsProposalsProposalSpace.About, and is useful for accessing the field via an interface.
func (v *ProposalsProposalsProposalSpace) GetAbout() string { return v.About }

// ProposalsResponse is returned by Proposals on success.
type ProposalsResponse struct {
	Proposals []ProposalsProposalsProposal `json:"proposals"`
}

// GetProposals returns ProposalsResponse.Proposals, and is useful for accessing the field via an interface.
func (v *ProposalsResponse) GetProposals() []ProposalsProposalsProposal { return v.Proposals }

// VoteResponse is returned by Vote on success.
type VoteResponse struct {
	Votes []VoteVotesVote `json:"votes"`
}

// GetVotes returns VoteResponse.Votes, and is useful for accessing the field via an interface.
func (v *VoteResponse) GetVotes() []VoteVotesVote { return v.Votes }

// VoteVotesVote includes the requested fields of the GraphQL type Vote.
type VoteVotesVote struct {
	Id      string  `json:"id"`
	Voter   string  `json:"voter"`
	Created int     `json:"created"`
	Choice  any     `json:"choice"`
	Vp      float64 `json:"vp"`
	Reason  string  `json:"reason"`
}

// GetId returns VoteVotesVote.Id, and is useful for accessing the field via an interface.
func (v *VoteVotesVote) GetId() string { return v.Id }

// GetVoter returns VoteVotesVote.Voter, and is useful for accessing the field via an interface.
func (v *VoteVotesVote) GetVoter() string { return v.Voter }

// GetCreated returns VoteVotesVote.Created, and is useful for accessing the field via an interface.
func (v *VoteVotesVote) GetCreated() int { return v.Created }

// GetChoice returns VoteVotesVote.Choice, and is useful for accessing the field via an interface.
func (v *VoteVotesVote) GetChoice() any { return v.Choice }

// GetVp returns VoteVotesVote.Vp, and is useful for accessing the field via an interface.
func (v *VoteVotesVote) GetVp() float64 { return v.Vp }

// GetReason returns VoteVotesVote.Reason, and is useful for accessing the field via an interface.
func (v *VoteVotesVote) GetReason() string { return v.Reason }

// VotesResponse is returned by Votes on success.
type VotesResponse struct {
	Votes []VotesVotesVote `json:"votes"`
}

// GetVotes returns VotesResponse.Votes, and is useful for accessing the field via an interface.
func (v *VotesResponse) GetVotes() []VotesVotesVote { return v.Votes }

// VotesVotesVote includes the requested fields of the GraphQL type Vote.
type VotesVotesVote struct {
	Id      string  `json:"id"`
	Voter   string  `json:"voter"`
	Created int     `json:"created"`
	Choice  any     `json:"choice"`
	Vp      float64 `json:"vp"`
	Reason  string  `json:"reason"`
}

// GetId returns VotesVotesVote.Id, and is useful for accessing the field via an interface.
func (v *VotesVotesVote) GetId() string { return v.Id }

// GetVoter returns VotesVotesVote.Voter, and is useful for accessing the field via an interface.
func (v *VotesVotesVote) GetVoter() string { return v.Voter }

// GetCreated returns VotesVotesVote.Created, and is useful for accessing the field via an interface.
func (v *VotesVotesVote) GetCreated() int { return v.Created }

// GetChoice returns VotesVotesVote.Choice, and is useful for accessing the field via an interface.
func (v *VotesVotesVote) GetChoice() any { return v.Choice }

// GetVp returns VotesVotesVote.Vp, and is useful for accessing the field via an interface.
func (v *VotesVotesVote) GetVp() float64 { return v.Vp }

// GetReason returns VotesVotesVote.Reason, and is useful for accessing the field via an interface.
func (v *VotesVotesVote) GetReason() string { return v.Reason }

// __ProposalInput is used internally by genqlient
type __ProposalInput struct {
	Id string `json:"id"`
}

// GetId returns __ProposalInput.Id, and is useful for accessing the field via an interface.
func (v *__ProposalInput) GetId() string { return v.Id }

// __ProposalsInput is used internally by genqlient
type __ProposalsInput struct {
	Created int `json:"created"`
}

// GetCreated returns __ProposalsInput.Created, and is useful for accessing the field via an interface.
func (v *__ProposalsInput) GetCreated() int { return v.Created }

// __VoteInput is used internally by genqlient
type __VoteInput struct {
	Id string `json:"id"`
}

// GetId returns __VoteInput.Id, and is useful for accessing the field via an interface.
func (v *__VoteInput) GetId() string { return v.Id }

// __VotesInput is used internally by genqlient
type __VotesInput struct {
	Proposal string `json:"proposal"`
	Created  int    `json:"created"`
}

// GetProposal returns __VotesInput.Proposal, and is useful for accessing the field via an interface.
func (v *__VotesInput) GetProposal() string { return v.Proposal }

// GetCreated returns __VotesInput.Created, and is useful for accessing the field via an interface.
func (v *__VotesInput) GetCreated() int { return v.Created }

// The query or mutation executed by Proposal.
const Proposal_Operation = `
query Proposal ($id: String) {
	proposals(where: {id:$id}) {
		id
		link
		title
		body
		choices
		start
		end
		state
		author
		created
		space {
			id
			name
			about
		}
	}
}
`

func Proposal(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*ProposalResponse, error) {
	req := &graphql.Request{
		OpName: "Proposal",
		Query:  Proposal_Operation,
		Variables: &__ProposalInput{
			Id: id,
		},
	}
	var err error

	var data ProposalResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by Proposals.
const Proposals_Operation = `
query Proposals ($created: Int) {
	proposals(first: 1000, where: {created_gte:$created}, orderBy: "created", orderDirection: asc) {
		id
		link
		title
		body
		choices
		start
		end
		state
		author
		created
		space {
			id
			name
			about
		}
	}
}
`

func Proposals(
	ctx context.Context,
	client graphql.Client,
	created int,
) (*ProposalsResponse, error) {
	req := &graphql.Request{
		OpName: "Proposals",
		Query:  Proposals_Operation,
		Variables: &__ProposalsInput{
			Created: created,
		},
	}
	var err error

	var data ProposalsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by Vote.
const Vote_Operation = `
query Vote ($id: String) {
	votes(where: {id:$id}) {
		id
		voter
		created
		choice
		vp
		reason
	}
}
`

func Vote(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*VoteResponse, error) {
	req := &graphql.Request{
		OpName: "Vote",
		Query:  Vote_Operation,
		Variables: &__VoteInput{
			Id: id,
		},
	}
	var err error

	var data VoteResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by Votes.
const Votes_Operation = `
query Votes ($proposal: String, $created: Int) {
	votes(first: 1000, where: {proposal:$proposal,created_gte:$created}, orderBy: "created", orderDirection: asc) {
		id
		voter
		created
		choice
		vp
		reason
	}
}
`

func Votes(
	ctx context.Context,
	client graphql.Client,
	proposal string,
	created int,
) (*VotesResponse, error) {
	req := &graphql.Request{
		OpName: "Votes",
		Query:  Votes_Operation,
		Variables: &__VotesInput{
			Proposal: proposal,
			Created:  created,
		},
	}
	var err error

	var data VotesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
