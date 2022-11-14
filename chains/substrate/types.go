// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"math/big"
    "fmt"
	"github.com/uinb/chainbridge-utils/msg"
	"github.com/uinb/go-substrate-rpc-client/v4/scale"
	"github.com/uinb/go-substrate-rpc-client/v4/types"
)

type voteState struct {
	VotesFor     []types.AccountID
	VotesAgainst []types.AccountID
	Status       voteStatus
}

type voteStatus struct {
	IsActive   bool
	IsApproved bool
	IsRejected bool
}

func (m *voteStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsActive = true
	} else if b == 1 {
		m.IsApproved = true
	} else if b == 2 {
		m.IsRejected = true
	}

	return nil
}

// proposal represents an on-chain proposal
type proposal struct {
	depositNonce types.U64
	call         types.Call
	sourceId     types.U16
	resourceId   types.Bytes32
	method       string
}

// encode takes only nonce and call and encodes them for storage queries
func (p *proposal) encode() ([]byte, error) {
	return types.Encode(struct {
		types.U64
		types.Call
	}{p.depositNonce, p.call})
}

func (w *writer) createFungibleProposal(m msg.Message) (*proposal, error) {
	bigAmt := big.NewInt(0).SetBytes(m.Payload[0].([]byte))
	amount := types.NewU128(*bigAmt)
	recipient := types.NewAccountID(m.Payload[1].([]byte))
	depositNonce := types.U64(m.DepositNonce)

	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)
	if err != nil {
		return nil, err
	}
	call, err := types.NewCall(
		&meta,
		method,
		recipient,
		amount,
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := types.Encode(m.ResourceId)
		if err != nil {
			return nil, err
		}
		call.Args = append(call.Args, eRID...)
	}
/*    sourceU8 := types.U8(m.Source)
	bsource, err  := sourceU8.MarshalJSON()
	if err != nil {
	    return nil, err
	}
	sId := types.U16(0)
    err = sId.UnmarshalJSON(bsource)
	if err != nil {
	    return nil, err
	}*/
	return &proposal{
		depositNonce: depositNonce,
		call:         call,
		sourceId:     types.U16(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       method,
	}, nil
}

func (w *writer) createNonFungibleProposal(m msg.Message) (*proposal, error) {
	tokenId := types.NewU256(*big.NewInt(0).SetBytes(m.Payload[0].([]byte)))
	recipient := types.NewAccountID(m.Payload[1].([]byte))
	metadata := types.Bytes(m.Payload[2].([]byte))
	depositNonce := types.U64(m.DepositNonce)

	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)
	if err != nil {
		return nil, err
	}

	call, err := types.NewCall(
		&meta,
		method,
		recipient,
		tokenId,
		metadata,
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := types.Encode(m.ResourceId)
		if err != nil {
			return nil, err
		}
		call.Args = append(call.Args, eRID...)
	}

	return &proposal{
		depositNonce: depositNonce,
		call:         call,
		sourceId:     types.U16(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       method,
	}, nil
}

type H1024 [128]byte

// NewH1024 creates a new H1024 type
func NewH1024(b []byte) H1024 {

    h := H1024{}
    copy(h[:], b)
    return h
}

// Hex returns a hex string representation of the value (not of the encoded value)
func (h H1024) Hex() string {
    return fmt.Sprintf("%#x", h[:])
}

func (w *writer) createGenericProposal(m msg.Message) (*proposal, error) {
	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)
	if err != nil {
		return nil, err
	}

    w.log.Info("-------------------")
    w.log.Info(fmt.Sprintf("%#x", m.Payload[0].([]byte)[:]))
/*
    message, err := types.Encode(m.Payload[0].([]byte)[:])
    if err != nil {
        return nil, err
    } */
	call, err := types.NewCall(
		&meta,
		method,
		m.Payload[0].([]byte)[:],
	)
	if err != nil {
		return nil, err
	}

	dep, err := types.Encode(m.Depositer)
	if err != nil {
	    return nil, err
	}
	call.Args = append(call.Args, dep...)

	if w.extendCall {
		eRID, err := types.Encode(m.ResourceId)
		if err != nil {
			return nil, err
		}

		call.Args = append(call.Args, eRID...)
	}
	return &proposal{
		depositNonce: types.U64(m.DepositNonce),
		call:         call,
		sourceId:     types.U16(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       method,
	}, nil
}
