package messages

import (
	"errors"
	"log"
)

type CHANNEL_STATUS byte

const (
	CHANNEL_STATUS_UNASSIGNED CHANNEL_STATUS = 0
	CHANNEL_STATUS_ASSIGNED   CHANNEL_STATUS = 1
	CHANNEL_STATUS_SEARCHING  CHANNEL_STATUS = 2
	CHANNEL_STATUS_TRACKING   CHANNEL_STATUS = 3
)

const (
	CHANNEL_STATUS_MSG_ID   = byte(0x52)
	CHAMMEN_STATUS_MSG_NAME = "CHANNEL_STATUS"
)

type ChannelStatusReply struct {
	AntCommandInfo
	Status CHANNEL_STATUS
}

func CreateChannelStatusReply(msg *AntCommandMessage) (*ChannelStatusReply, error) {
	if msg.Id != CHANNEL_STATUS_MSG_ID {
		log.Println("invalid message: wrong message id")
		log.Printf("expected: %v, actual: %v", CHANNEL_STATUS_MSG_ID, msg.Id)

		return nil, errors.New("invallid message: wrong message id")
	}

	return &ChannelStatusReply{
		AntCommandInfo: newAntCommandInfo(msg.Id, msg.Name),
		Status:         CHANNEL_STATUS(msg.Data[0]),
	}, nil
}
