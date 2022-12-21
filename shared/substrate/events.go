// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/uinb/go-substrate-rpc-client/v4/types"
	events "github.com/uinb/ChainBridge/common"
)

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type Events struct {
	events.Events

}
