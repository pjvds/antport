package messages

import (
	"errors"
	"log"
)

const (
	RecvBroadcastDataCommandId   = byte(0x4e)
	RecvBroadcastDataCommandName = "RECV_BROADCAST_DATA"
)

type RecvBroadcastDataReply struct {
	AntCommandInfo

	ChannelNumber byte
	Data          []byte
}

func CreateRecvBroadcastDataReply(msg AntCommandMessage) (*RecvBroadcastDataReply, error) {
	if msg.Id != RecvBroadcastDataCommandId {
		log.Println("invalid message: wrong message id")
		log.Printf("expected: %v, actual: %v", RecvBroadcastDataCommandId, msg.Id)

		return nil, errors.New("invallid message: wrong message id")
	}

	return &RecvBroadcastDataReply{
		AntCommandInfo: newAntCommandInfo(msg.Id, msg.Name),
		Data:           msg.Data,
	}, nil
}
