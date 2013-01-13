package ant

import (
	"github.com/pjvds/antport/messages"
)

type AntNetwork struct {
	ant *AntContext

	// The Network Number is an 8-bit field with the 
	// range of acceptable values being from 0 to the 
	// maximum number defined by the ANT implementation.
	Number byte

	// The Network Key is an 8-byte field which is configurable
	// by the host application. A particular Network
	// Number will have a corresponding Network Key.  
	// The Network Number and the Network Key together provide 
	// the ability to deploy a network with varied levels of 
	// access control and security options. 
	Key [8]byte
}

func (network AntNetwork) SetNetworkKey(key [8]byte) {
	ant := network.ant
	cmd := messages.CreateSetNetworkKeyCommand(network.Number, key)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}
