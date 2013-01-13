package messages

import (
	"errors"
	"log"
)

const (
	CAPABILITIES_MSG_ID   = byte(0x54)
	CapabilitiesCommandName = "CAPABILITIES"
)

type CapabilitiesReply struct {
	AntCommandInfo
	MaxChannels     byte
	MaxNetworks     byte
	StandardOptions byte
}

func NewCapabilitiesReply(msg *AntCommandMessage) (*CapabilitiesReply, error) {
	if msg.Id != CAPABILITIES_MSG_ID {
		log.Println("invalid message: wrong message id")
		log.Printf("expected: %v, actual: %v", CAPABILITIES_MSG_ID, msg.Id)

		return nil, errors.New("invallid message: wrong message id")
	}

	return &CapabilitiesReply{
		AntCommandInfo:  newAntCommandInfo(msg.Id, msg.Name),
		MaxChannels:     msg.Data[0],
		MaxNetworks:     msg.Data[1],
		StandardOptions: msg.Data[2],
	}, nil
}
