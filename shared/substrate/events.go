// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/uinb/go-substrate-rpc-client/v4/types"
)

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Transferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Burned struct {
	Phase   types.Phase
	TokenId types.AccountID
	Topics  []types.Hash
}

type EventExampleRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventChainBridgeTransferRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.U32
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventRadClaimsClaimed is emitted when RAD Tokens have been claimed
type EventRadClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventRadClaimsRootHashStored is emitted when RootHash has been stored for the correspondent RAD Claims batch
type EventRadClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventNftTransferred is emitted when the ownership of the asset has been transferred to the account
type EventNftTransferred struct {
	Phase      types.Phase
	RegistryId RegistryId
	AssetId    AssetId
	Who        types.AccountID
	Topics     []types.Hash
}

// EventRegistryMint is emitted when successfully minting an NFT
type EventRegistryMint struct {
	Phase      types.Phase
	RegistryId RegistryId
	TokenId    TokenId
	Topics     []types.Hash
}

// EventRegistryRegistryCreated is emitted when successfully creating a NFT registry
type EventRegistryRegistryCreated struct {
	Phase      types.Phase
	RegistryId RegistryId
	Topics     []types.Hash
}

// EventRegistryTmp is emitted only for testing
type EventRegistryTmp struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// lpos event
type EventPlanNewEra struct {
	Phase    types.Phase
	EraIndex types.U32
	Topics   []types.Hash
}
type EventPlanNewEraFailed struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventTriggerNewEra struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EraIndex types.U32

type EventEraPayout struct {
	Phase              types.Phase
	EraIndex           EraIndex
	ExcludedValidators []types.AccountID
	Topics             []types.Hash
}
type EventEraPayoutFailed struct {
	Phase    types.Phase
	EraIndex EraIndex
	Topics   []types.Hash
}
type SessionIndex types.U32

type EventOldSlashingReportDiscarded struct {
	Phase        types.Phase
	SessionIndex SessionIndex
	Topics       []types.Hash
}

//appchain event
type Validator struct {
	ValidatorIdInAppchain types.AccountID
	TotalStake            types.U128
}
type EventNewPlannedValidators struct {
	Phase      types.Phase
	SetId      types.U32
	Validators []Validator
	Topics     []types.Hash
}

type EventLocked struct {
	Phase    types.Phase
	Sender   types.AccountID
	Receiver types.Bytes
	Amount   types.U128
	Sequence types.U64
	Topics   []types.Hash
}
type EventUnlocked struct {
	Phase    types.Phase
	Sender   types.Bytes
	Receiver types.AccountID
	Amount   types.U128
	Topics   []types.Hash
}
type EventUnlockFailed struct {
	Phase    types.Phase
	Sender   types.Bytes
	Receiver types.AccountID
	Amount   types.U128
	Topics   []types.Hash
}
type EventAssetMinted struct {
	Phase    types.Phase
	AssetId  types.U32
	Sender   types.Bytes
	Receiver types.AccountID
	Amount   types.U128
	Sequence types.OptionU32
	Topics   []types.Hash
}

type EventAssetBurned struct {
	Phase    types.Phase
	AssetId  types.U32
	Sender   types.AccountID
	Receiver types.Bytes
	Amount   types.U128
	Sequence types.U64
	Topics   []types.Hash
}

type EventAssetMintFailed struct {
	Phase    types.Phase
	AssetId  types.U32
	Sender   types.Bytes
	Receiver types.AccountID
	Amount   types.U128
	Sequence types.OptionU32
	Topics   []types.Hash
}

type EventAssetIdGetFailed struct {
	Phase    types.Phase
	TokenId  types.Bytes
	Sender   types.Bytes
	Receiver types.AccountID
	Amount   types.U128
	Topics   []types.Hash
}

type EventTransferredFromPallet struct {
	Phase    types.Phase
	Receiver types.AccountID
	Amount   types.U128
	Topics   []types.Hash
}

type EventNftLocked struct {
	Phase    types.Phase
	Sender   types.AccountID
	Receiver types.Bytes
	Class    types.U128
	Instance types.U128
	Sequence types.U64
	Topics   []types.Hash
}

type EventNftUnLocked struct {
	Phase    types.Phase
	Sender   types.Bytes
	Receiver types.AccountID
	Class    types.U128
	Instance types.U128
	Topics   []types.Hash
}

type EventNftUnlockFailed struct {
	Phase    types.Phase
	Sender   types.Bytes
	Receiver types.AccountID
	Class    types.U128
	Instance types.U128
	Topics   []types.Hash
}

type EventForceAssetMinted struct {
	Phase   types.Phase
	AssetId types.U32
	Who     types.AccountID
	Amount  types.U128
	Topics  []types.Hash
}

type Events struct {
	types.EventRecords
	ChainBridgeEvents

}
