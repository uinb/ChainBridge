package substrate

import (
	"fmt"
	gethrpc "github.com/centrifuge/go-substrate-rpc-client/gethrpc"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"sync"
)

type ExtrinsicStatusSubscription struct {
	sub      *gethrpc.ClientSubscription
	channel  chan types.ExtrinsicStatus
	quitOnce sync.Once // ensures quit is closed once
}

func (s *ExtrinsicStatusSubscription) Unsubscribe() {
	s.sub.Unsubscribe()
	s.quitOnce.Do(func() {
		close(s.channel)
	})
}


// Chan returns the subscription channel.
//
// The channel is closed when Unsubscribe is called on the subscription.
func (s *ExtrinsicStatusSubscription) Chan() <-chan types.ExtrinsicStatus {
	return s.channel
}

// Err returns the subscription error channel. The intended use of Err is to schedule
// resubscription when the client connection is closed unexpectedly.
//
// The error channel receives a value when the subscription has ended due
// to an error. The received error is nil if Close has been called
// on the underlying client and no other error has occurred.
//
// The error channel is closed when Unsubscribe is called on the subscription.
func (s *ExtrinsicStatusSubscription) Err() <-chan error {
	return s.sub.Err()
}



func (c *Connection) watchSubmission1(sub *ExtrinsicStatusSubscription) error {
	for {
		select {
		case <-c.stop:
			return TerminatedError
		case status := <-sub.Chan():
			switch {
			case status.IsInBlock:
				c.log.Trace("Extrinsic included in block", "block", status.AsInBlock.Hex())
				return nil
			case status.IsRetracted:
				return fmt.Errorf("extrinsic retracted: %s", status.AsRetracted.Hex())
			case status.IsDropped:
				return fmt.Errorf("extrinsic dropped from network")
			case status.IsInvalid:
				return fmt.Errorf("extrinsic invalid")
			}
		case err := <-sub.Err():
			c.log.Trace("Extrinsic subscription error", "err", err)
			return err
		}
	}
}

