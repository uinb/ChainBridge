// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	centEvents "github.com/centrifuge/chain-custom-types"
	events "github.com/centrifuge/chainbridge-substrate-events"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
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

type cEvents = centEvents.Events
type Events struct {
	types.EventRecords
	events.Events
	FusotaoEvents
	cEvents
	Erc721_Minted            []EventErc721Minted            //nolint:stylecheck,golint
	Erc721_Transferred       []EventErc721Transferred       //nolint:stylecheck,golint
	Erc721_Burned            []EventErc721Burned            //nolint:stylecheck,golint
	Example_Remark           []EventExampleRemark           //nolint:stylecheck,golint
	Registry_Mint            []EventRegistryMint            //nolint:stylecheck,golint
	Registry_RegistryCreated []EventRegistryRegistryCreated //nolint:stylecheck,golint
	Registry_RegistryTmp     []EventRegistryTmp             //nolint:stylecheck,golint
	// add octopus pallets events
	ChainBridgeTransfer_Remark            []EventChainBridgeTransferRemark      //nolint:stylecheck,golint
	ChainBridgeAssets_Created             []types.EventAssetCreated             //nolint:stylecheck,golint
	ChainBridgeAssets_Issued              []types.EventAssetIssued              //nolint:stylecheck,golint
	ChainBridgeAssets_Transferred         []types.EventAssetTransferred         //nolint:stylecheck,golint
	ChainBridgeAssets_Burned              []types.EventAssetBurned              //nolint:stylecheck,golint
	ChainBridgeAssets_TeamChanged         []types.EventAssetTeamChanged         //nolint:stylecheck,golint
	ChainBridgeAssets_OwnerChanged        []types.EventAssetOwnerChanged        //nolint:stylecheck,golint
	ChainBridgeAssets_Frozen              []types.EventAssetFrozen              //nolint:stylecheck,golint
	ChainBridgeAssets_Thawed              []types.EventAssetThawed              //nolint:stylecheck,golint
	ChainBridgeAssets_AssetFrozen         []types.EventAssetAssetFrozen         //nolint:stylecheck,golint
	ChainBridgeAssets_AssetThawed         []types.EventAssetAssetThawed         //nolint:stylecheck,golint
	ChainBridgeAssets_Destroyed           []types.EventAssetDestroyed           //nolint:stylecheck,golint
	ChainBridgeAssets_ForceCreated        []types.EventAssetForceCreated        //nolint:stylecheck,golint
	ChainBridgeAssets_MetadataSet         []types.EventAssetMetadataSet         //nolint:stylecheck,golint
	ChainBridgeAssets_MetadataCleared     []types.EventAssetMetadataCleared     //nolint:stylecheck,golint
	ChainBridgeAssets_ApprovedTransfer    []types.EventAssetApprovedTransfer    //nolint:stylecheck,golint
	ChainBridgeAssets_ApprovalCancelled   []types.EventAssetApprovalCancelled   //nolint:stylecheck,golint
	ChainBridgeAssets_TransferredApproved []types.EventAssetTransferredApproved //nolint:stylecheck,golint
	ChainBridgeAssets_AssetStatusChanged  []types.EventAssetAssetStatusChanged  //nolint:stylecheck,golint

	OctopusLpos_PlanNewEra                 []EventPlanNewEra                 //nolint:stylecheck,golint
	OctopusLpos_PlanNewEraFailed           []EventPlanNewEraFailed           //nolint:stylecheck,golint
	OctopusLpos_TriggerNewEra              []EventTriggerNewEra              //nolint:stylecheck,golint
	OctopusLpos_EraPayout                  []EventEraPayout                  //nolint:stylecheck,golint
	OctopusLpos_EraPayoutFailed            []EventEraPayoutFailed            //nolint:stylecheck,golint
	OctopusLpos_OldSlashingReportDiscarded []EventOldSlashingReportDiscarded //nolint:stylecheck,golint

	OctopusAppchain_NewPlannedValidators  []EventNewPlannedValidators  //nolint:stylecheck,golint
	OctopusAppchain_Locked                []EventLocked                //nolint:stylecheck,golint
	OctopusAppchain_Unlocked              []EventUnlocked              //nolint:stylecheck,golint
	OctopusAppchain_UnlockFailed          []EventUnlockFailed          //nolint:stylecheck,golint
	OctopusAppchain_AssetMinted           []EventAssetMinted           //nolint:stylecheck,golint
	OctopusAppchain_AssetBurned           []EventAssetBurned           //nolint:stylecheck,golint
	OctopusAppchain_AssetMintFailed       []EventAssetMintFailed       //nolint:stylecheck,golint
	OctopusAppchain_AssetIdGetFailed      []EventAssetIdGetFailed      //nolint:stylecheck,golint
	OctopusAppchain_TransferredFromPallet []EventTransferredFromPallet //nolint:stylecheck,golint
	OctopusAppchain_NftLocked             []EventNftLocked             //nolint:stylecheck,golint
	OctopusAppchain_NftUnlocked           []EventNftUnLocked           //nolint:stylecheck,golint
	OctopusAppchain_NftUnlockFailed       []EventNftUnlockFailed       //nolint:stylecheck,golint
	OctopusAppchain_ForceAssetMinted      []EventForceAssetMinted      //nolint:stylecheck,golint
}
