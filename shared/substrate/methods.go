// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = BridgePalletName + ".add_relayer"
var SetResourceMethod Method = BridgePalletName + ".set_resource"
var SetThresholdMethod Method = BridgePalletName + ".set_threshold"
var WhitelistChainMethod Method = BridgePalletName + ".whitelist_chain"

var ExampleTransferNativeMethod Method = "Example.transfer_native"
var ExampleTransferErc721Method Method = "Example.transfer_erc721"
var ExampleTransferHashMethod Method = "Example.transfer_hash"
var ExampleMintErc721Method Method = "Example.mint_erc721"
var ExampleTransferMethod Method = "Example.transfer"
var ExampleRemarkMethod Method = "Example.remark"

var ChainBridgeTransferNativeMethod Method = "ChainBridgeTransfer.transfer_native"
var ChainBridgeTransferErc721Method Method = "ChainBridgeTransfer.transfer_erc721"
var ChainBridgeTransferHashMethod Method = "ChainBridgeTransfer.transfer_hash"
var ChainBridgeTransferMintErc721Method Method = "ChainBridgeTransfer.mint_erc721"
var ChainBridgeTransferMethod Method = "ChainBridgeTransfer.transfer"
var ChainBridgeTransferRemarkMethod Method = "ChainBridgeTransfer.remark"

var Erc721MintMethod Method = "Erc721.mint"
var SudoMethod Method = "Sudo.sudo"
