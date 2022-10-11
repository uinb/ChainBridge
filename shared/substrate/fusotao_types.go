// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
    "github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type FusotaoEvents struct {

    Foundation_PreLockedFundUnlocked    []EventPreLockedFundUnlocked
    Foundation_UnlockedFundAllBalance   []EventUnlockedFundAllBalance
    Reward_RewardClaimed                []EventRewardClaimed
    Reward_EraRewardsUpdated            []EventEraRewardsUpdated
    Agent_ControllerTxCompleted         []EventControllerTxCompleted

    Token_TokenIssued                   []EventTokenIssued
    Token_TokenTransfered               []EventTokenTransfered
   	Token_TokenReserved                 []EventTokenReserved
   	Token_TokenUnreserved               []EventTokenUnreserved
   	Token_TokenMinted                   []EventTokenMinted
   	Token_TokenBurned                   []EventTokenBurned
   	Token_TokenRepatriated              []EventTokenRepatriated

   	Verifier_TokenHosted                []EventTokenHosted
   	Verifier_TokenRevoked               []EventTokenRevoked
   	Verifier_TaoStaked                  []EventTaoStaked
   	Verifier_TaoUnstaked                []EventTaoUnstaked
   	Verifier_TaoUnstakeUnlock           []EventTaoUnstakeUnlock
   	Verifier_DominatorClaimed           []EventDominatorClaimed
   	Verifier_DominatorOnline            []EventDominatorOnline
   	Verifier_DominatorOffline           []EventDominatorOffline
   	Verifier_DominatorSlashed           []EventDominatorSlashed
   	Verifier_DominatorEvicted           []EventDominatorEvicted
   	Verifier_DominatorInactive          []EventDominatorInActive

}

type EventPreLockedFundUnlocked struct {
    Phase      types.Phase
    Who        types.AccountID
    Balance     types.U128
    Topics  []types.Hash
}
type EventUnlockedFundAllBalance struct {
    Phase      types.Phase
    Who        types.AccountID
    Topics  []types.Hash
}
type EventRewardClaimed struct {
    Phase      types.Phase
    Who      types.AccountID
    Balance  types.U128
    Topics   []types.Hash
}
type EventEraRewardsUpdated struct {
    Phase      types.Phase
    Balance     types.U128
    Topics  []types.Hash
}
type EventControllerTxCompleted struct {
    Phase      types.Phase
    Topics  []types.Hash
}

type EventTokenHosted struct {
    Phase      types.Phase
    Who        types.AccountID
	Who1       types.AccountID
	TokenId    types.U32
	Amount     types.U128
    Topics  []types.Hash
}
type EventTokenRevoked struct {
    Phase      types.Phase
    Who        types.AccountID
	Who1       types.AccountID
	TokenId    types.U32
	Amount     types.U128
    Topics  []types.Hash
}
type EventTaoStaked struct {
    Phase      types.Phase
    Who        types.AccountID
	Who1       types.AccountID
	Amount     types.U128
    Topics  []types.Hash
}
type EventTaoUnstaked struct {
    Phase      types.Phase
    Who        types.AccountID
	Who1       types.AccountID
	Amount     types.U128
    Topics  []types.Hash
}
type EventTaoUnstakeUnlock struct {
    Phase      types.Phase
	Who       types.AccountID
	Amount     types.U128
    Topics  []types.Hash
}
type EventDominatorClaimed struct {
    Phase      types.Phase
	Who        types.AccountID
    Topics     []types.Hash
}
type EventDominatorOnline struct {
    Phase      types.Phase
	Who        types.AccountID
    Topics     []types.Hash
}
type EventDominatorOffline struct {
    Phase      types.Phase
	Who        types.AccountID
    Topics     []types.Hash
}
type EventDominatorSlashed struct {
    Phase      types.Phase
	Who        types.AccountID
    Topics     []types.Hash
}
type EventDominatorEvicted struct {
    Phase      types.Phase
	Who        types.AccountID
    Topics     []types.Hash
}
type EventDominatorInActive struct {
    Phase      types.Phase
	Who        types.AccountID
    Topics     []types.Hash
}

type EventTokenIssued struct {
    Phase      types.Phase
    TokenId    types.U32
    Contract   types.Bytes
    Name       types.Bytes
    Topics  []types.Hash
}

type EventTokenTransfered struct {
    Phase      types.Phase
    TokenId     types.U32
    Who         types.AccountID
    Whoo        types.AccountID
    Balance     types.U128
    Topics      []types.Hash
}

type EventTokenReserved struct {
    Phase      types.Phase
    TokenId     types.U32
    Who         types.AccountID
    Balance     types.U128
    Topics  []types.Hash
}

type EventTokenUnreserved struct {
    Phase      types.Phase
    TokenId     types.U32
    Who         types.AccountID
    Balance     types.U128
    Topics  []types.Hash
}

type EventTokenMinted struct {
    Phase      types.Phase
    TokenId     types.U32
    Who         types.AccountID
    Balance     types.U128
    Topics  []types.Hash
}

type EventTokenBurned struct {
    Phase      types.Phase
    TokenId     types.U32
    Who         types.AccountID
    Balance     types.U128
    Topics  []types.Hash
}


type EventTokenRepatriated struct {
    Phase      types.Phase
    TokenId     types.U32
    Who         types.AccountID
    Who1         types.AccountID
    Balance     types.U128
    Topics  []types.Hash
}

