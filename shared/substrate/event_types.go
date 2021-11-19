package utils


import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type SubstrateEvents struct {
	ChainBridge_FungibleTransfer        []EventFungibleTransfer        //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventNonFungibleTransfer     //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventGenericTransfer         //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventRelayerThresholdChanged //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainWhitelisted        //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventRelayerAdded            //nolint:stylecheck,golint
	ChainBridge_RelayerRemoved          []EventRelayerRemoved          //nolint:stylecheck,golint
	ChainBridge_VoteFor                 []EventVoteFor                 //nolint:stylecheck,golint
	ChainBridge_VoteAgainst             []EventVoteAgainst             //nolint:stylecheck,golint
	ChainBridge_ProposalApproved        []EventProposalApproved        //nolint:stylecheck,golint
	ChainBridge_ProposalRejected        []EventProposalRejected        //nolint:stylecheck,golint
	ChainBridge_ProposalSucceeded       []EventProposalSucceeded       //nolint:stylecheck,golint
	ChainBridge_ProposalFailed          []EventProposalFailed          //nolint:stylecheck,golint

	Token_TokenIssued                   []EventTokenIssued
	Token_TokenTransfered               []EventTokenTransfered
	Token_TokenReserved                 []EventTokenReserved
	Token_TokenUnreserved               []EventTokenUnreserved
	Token_TokenBurned                   []EventTokenBurned
	Token_TokenRepatriated              []EventTokenRepatriated

	Receipts_DominatorClaimed           []ReceiptsDominatorClaimed
	Receipts_CoinHosted                 []ReceiptsCoinHosted
	Receipts_TokenHosted                []ReceiptsTokenHosted
	Receipts_CoinRevoked                []ReceiptsCoinRevoked
	Receipts_TokenRevoked               []ReceiptsTokenRevoked

	Council_CandidateNominated          []CouncilCandidateNominated
	Council_MembersAnnounced            []CouncilMembersAnnounced

	Elections_StartProposal             []ElectionStartProposal
	Elections_StopProposal              []ElectionStopProposal
	Elections_Voted                     []ElectionVoted
	Elections_AddCandidate              []ElectionAddCandidate
}

type ElectionStartProposal struct {
	Phase types.Phase
	BlockNumber1   types.U32
	BlockNumber2   types.U32
	VoteIndex      types.U32
	Topics         []types.Hash

}

type ElectionAddCandidate struct {
	Phase types.Phase
	Account1   types.AccountID
	Account2   types.AccountID
	Topics         []types.Hash

}

type ElectionVoted struct {
	Phase types.Phase
	Account1   types.AccountID
	Account2   types.AccountID
	Balance    types.U128
	Topics         []types.Hash

}


type ElectionStopProposal struct {
	Phase types.Phase
	VoteIndex      types.U32
	Topics         []types.Hash

}


type CouncilCandidateNominated struct {
	Phase      types.Phase
	AccountId  types.AccountID
	Topics     []types.Hash
}


type CouncilMembersAnnounced struct {
	Phase      types.Phase
	tokenId    types.U32
	Topics     []types.Hash
}

type ReceiptsDominatorClaimed struct {
	Phase      types.Phase

	AccountId 	   types.AccountID
	AmountCoin types.U128
	Topics     []types.Hash
}
type ReceiptsCoinHosted struct {
	Phase      types.Phase

	From	   types.AccountID
	To   	   types.AccountID
	AmountCoin types.U128
	Topics     []types.Hash
}



type ReceiptsTokenHosted struct {
	Phase      types.Phase

	From	   types.AccountID
	To   	   types.AccountID
	TokenId    types.U32
	AmountCoin types.U128
	Topics     []types.Hash
}




type ReceiptsCoinRevoked struct {
	Phase      types.Phase
	From	   types.AccountID
	To   	   types.AccountID
	AmountCoin types.U128
	Topics     []types.Hash
}


type ReceiptsTokenRevoked struct {
	Phase      types.Phase

	From	   types.AccountID
	To   	   types.AccountID
	TokenId    types.U32
	AmountCoin types.U128
	Topics     []types.Hash
}



type EventTokenTransfered struct {
	Phase      types.Phase
	TokenId    types.U32
	From 	   types.AccountID
	To   	   types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventTokenRepatriated struct {
	Phase      types.Phase
	TokenId    types.U32
	From 	   types.AccountID
	To   	   types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventTokenBurned struct {
	Phase     types.Phase
	TokenId   types.U32
	AccountId types.AccountID
	Balance   types.U128
	Topics     []types.Hash
}

type EventTokenUnreserved struct {
	Phase     types.Phase
	TokenId   types.U32
	AccountId types.AccountID
	Balance   types.U128
	Topics     []types.Hash
}

type EventTokenReserved struct {
	Phase     types.Phase
	TokenId   types.U32
	AccountId types.AccountID
	Balance   types.U128
	Topics     []types.Hash
}

type EventTokenIssued struct {
	Phase     types.Phase
	TokenId   types.U32
	AccountId types.AccountID
	Balance   types.U128
	Topics     []types.Hash
}
type EventFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Amount       types.U256
	Recipient    types.Bytes
	Topics       []types.Hash
}


type EventNonFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	TokenId      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventGenericTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventRelayerThresholdChanged struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventChainWhitelisted struct {
	Phase   types.Phase
	ChainId types.U8
	Topics  []types.Hash
}

type EventRelayerAdded struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventRelayerRemoved struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventVoteFor struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventVoteAgainst struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventProposalApproved struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalRejected struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalSucceeded struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalFailed struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}
