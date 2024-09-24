package nouns

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/nouns"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config               *config.Module
	ethereumClient       ethereum.Client
	tokenClient          token.Client
	nounsAuctionFilterer *nouns.NounsAuctionHouseFilterer
	nounsDAOFilterer     *nouns.NounsDAOFilterer
	nounsTokenFilterer   *nouns.NounsTokenFilterer
}

func (w *worker) Name() string {
	return decentralized.Nouns.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
	}
}

func (w *worker) Platform() string {
	return decentralized.PlatformNouns.String()
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Collectible,
		tag.Governance,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.CollectibleAuction,
		typex.CollectibleMint,
		typex.GovernanceProposal,
		typex.GovernanceVote,
	}
}

func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			nouns.AddressNouns,
			nouns.AddressNounsAuction,
			nouns.AddressNounsDAO,
		},
		LogTopics: []common.Hash{
			nouns.EventAuctionBid,
			nouns.EventAuctionSettled,
			nouns.EventAuctionCreated,
			nouns.EventNounCreated,
			nouns.EventNounsProposal,
			nouns.EventNounsVote,
		},
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*activityx.Action
			err     error
		)

		switch {
		case w.matchNounsAuctionBid(log):
			actions, err = w.handleNounsAuctionBid(ctx, ethereumTask, log)
			activity.Type = typex.CollectibleAuction
		case w.matchNounsAuctionSettled(log):
			actions, err = w.handleNounsAuctionSettled(ctx, ethereumTask, log)
			activity.Type = typex.CollectibleAuction
		case w.matchNounsAuctionCreated(log):
			actions, err = w.handleNounsAuctionCreated(ctx, ethereumTask, log)
			activity.Type = typex.CollectibleAuction
		case w.matchNounCreated(log):
			actions, err = w.handleNounCreated(ctx, ethereumTask, log)
			activity.Type = typex.CollectibleMint
		case w.matchNounsProposal(log):
			actions, err = w.handleNounsProposal(ctx, ethereumTask, log)
			activity.Type = typex.GovernanceProposal
		case w.matchNounsVote(log):
			actions, err = w.handleNounsVote(ctx, ethereumTask, log)
			activity.Type = typex.GovernanceVote
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		activity.Actions = append(activity.Actions, actions...)
	}

	// zap.L().Info("Processing task", zap.Any("task", ethereumTask))

	return activity, nil
}

func (w *worker) matchNounsAuctionBid(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, nouns.AddressNounsAuction) && contract.MatchEventHashes(log.Topics[0], nouns.EventAuctionBid)
}

func (w *worker) matchNounsAuctionSettled(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, nouns.AddressNounsAuction) && contract.MatchEventHashes(log.Topics[0], nouns.EventAuctionSettled)
}

func (w *worker) matchNounsAuctionCreated(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, nouns.AddressNounsAuction) && contract.MatchEventHashes(log.Topics[0], nouns.EventAuctionCreated)
}

func (w *worker) matchNounCreated(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, nouns.AddressNouns) && contract.MatchEventHashes(log.Topics[0], nouns.EventNounCreated)
}

func (w *worker) matchNounsProposal(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, nouns.AddressNounsDAO) && contract.MatchEventHashes(log.Topics[0], nouns.EventNounsProposal)
}

func (w *worker) matchNounsVote(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, nouns.AddressNounsDAO) && contract.MatchEventHashes(log.Topics[0], nouns.EventNounsVote)
}

func (w *worker) handleNounsAuctionBid(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.nounsAuctionFilterer.ParseAuctionBid(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse auction bid event: %w", err)
	}

	action, err := w.buildCollectibleAuctionAction(ctx, task, event.Sender, nouns.AddressNounsAuction, metadata.ActionCollectibleAuctionBid, event.NounId, big.NewInt(1), event.Value)

	if err != nil {
		return nil, fmt.Errorf("build collectible auction action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) handleNounsAuctionSettled(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.nounsAuctionFilterer.ParseAuctionSettled(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse auction settled event: %w", err)
	}

	action, err := w.buildCollectibleAuctionAction(ctx, task, nouns.AddressNounsAuction, event.Winner, metadata.ActionCollectibleAuctionFinalize, event.NounId, big.NewInt(1), event.Amount)
	if err != nil {
		return nil, fmt.Errorf("build collectible auction action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) handleNounsAuctionCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.nounsAuctionFilterer.ParseAuctionCreated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse auction created event: %w", err)
	}

	action, err := w.buildCollectibleAuctionAction(ctx, task, nouns.AddressNounsAuction, nouns.AddressNounsAuction, metadata.ActionCollectibleAuctionCreate, event.NounId, big.NewInt(1), nil)
	if err != nil {
		return nil, fmt.Errorf("build collectible auction action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) buildCollectibleAuctionAction(ctx context.Context, task *source.Task, sender, receiver common.Address, action metadata.CollectibleAuctionAction, nftID, nftValue, offerValue *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &(nouns.AddressNouns), nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %s: %w", &(nouns.AddressNouns), nftID, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(nftValue, 0))

	var offerTokenMetadata *metadata.Token

	if offerValue != nil {
		offerTokenMetadata, err = w.tokenClient.Lookup(ctx, task.ChainID, nil, nil, task.Header.Number)
		if err != nil {
			return nil, fmt.Errorf("lookup offer token metadata: %w", err)
		}

		offerTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(offerValue, 0))
	}

	return &activityx.Action{
		Type:     typex.CollectibleAuction,
		Platform: w.Platform(),
		From:     sender.String(),
		To:       receiver.String(),
		Metadata: metadata.CollectibleAuction{
			Action: action,
			Token:  *tokenMetadata,
			Cost:   offerTokenMetadata,
		},
	}, nil
}

func (w *worker) handleNounCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.nounsTokenFilterer.ParseNounCreated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse noun created event: %w", err)
	}

	action, err := w.buildCollectibleMintAction(ctx, task, ethereum.AddressGenesis, nouns.AddressNounsAuction, nouns.AddressNouns, event.TokenId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build collectible mint action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) buildCollectibleMintAction(ctx context.Context, task *source.Task, from, to, contract common.Address, id, value *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &contract, id, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %s: %w", contract, id, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	return &activityx.Action{
		Type:     typex.CollectibleMint,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleTransfer(*tokenMetadata),
	}, nil
}

func (w *worker) handleNounsProposal(_ context.Context, _ *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.nounsDAOFilterer.ParseProposalCreatedWithRequirements(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse proposal created event: %w", err)
	}

	action := w.buildGovernanceProposalAction(event.Proposer, nouns.AddressNounsDAO, event.Id, event.Description, event.StartBlock, event.EndBlock)

	return []*activityx.Action{action}, nil
}

func (w *worker) buildGovernanceProposalAction(from, to common.Address, id *big.Int, description string, start, end *big.Int) *activityx.Action {
	return &activityx.Action{
		Type:     typex.GovernanceProposal,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.GovernanceProposal{
			ID:         id.String(),
			Body:       description,
			Options:    []string{},
			StartBlock: decimal.NewFromBigInt(start, 0).String(),
			EndBlock:   decimal.NewFromBigInt(end, 0).String(),
		},
	}
}

func (w *worker) handleNounsVote(_ context.Context, _ *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.nounsDAOFilterer.ParseVoteCast(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse vote cast event: %w", err)
	}

	// Check for nil addresses and proposalID
	if event.Voter == (common.Address{}) {
		return nil, fmt.Errorf("voter address is nil")
	}

	if event.ProposalId == nil {
		return nil, fmt.Errorf("proposalID is nil")
	}

	// Set voteAction using the event.Support
	var voteAction metadata.GovernanceVoteAction

	switch event.Support {
	case 0:
		voteAction = metadata.ActionGovernanceVoteAgainst
	case 1:
		voteAction = metadata.ActionGovernanceVoteFor
	case 2:
		voteAction = metadata.ActionGovernanceVoteAbstain
	default:
		return nil, fmt.Errorf("invalid support value: %d", event.Support)
	}

	action := w.buildGovernanceVoteAction(event.Voter, nouns.AddressNounsDAO, event.ProposalId, event.Votes, event.Reason, voteAction)

	return []*activityx.Action{action}, nil
}

func (w *worker) buildGovernanceVoteAction(from, to common.Address, proposalID, votes *big.Int, reason string, action metadata.GovernanceVoteAction) *activityx.Action {
	return &activityx.Action{
		Type:     typex.GovernanceVote,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.GovernanceVote{
			Action: action,
			Count:  uint64(decimal.NewFromBigInt(votes, 0).IntPart()),
			Reason: reason,
			Proposal: metadata.GovernanceProposal{
				ID:   proposalID.String(),
				Link: fmt.Sprintf("https://nouns.wtf/vote/%v", proposalID),
			},
		},
	}
}

// NewWorker creates a new Nouns worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var (
		err      error
		instance = worker{
			config: config,
		}
	)

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize contract filterers.
	instance.nounsAuctionFilterer = lo.Must(nouns.NewNounsAuctionHouseFilterer(nouns.AddressNounsAuction, nil))
	instance.nounsDAOFilterer = lo.Must(nouns.NewNounsDAOFilterer(nouns.AddressNounsDAO, nil))
	instance.nounsTokenFilterer = lo.Must(nouns.NewNounsTokenFilterer(nouns.AddressNouns, nil))

	return &instance, nil
}
