package messages

import (
	"errors"
)

const (
	CHANNEL_RESPONSE_MSG_ID   = byte(0x40)
	CHANNEL_RESPONSE_MSG_NAME = "CHANNEL_RESPONSE"
)

/* The Response/Event Messages are messages sent from the ANT device to the controller device, either in 
response to a message (see Section 9.3 for a list of messages that generate responses), or as generated 
by an RF event on the ANT device. */
type ChannelResponse struct {
	AntCommandInfo

	// The channel number of the channel associated with the event. 
	ChannelNumber byte

	// ID of the message being responded too. Set to 1 if being sent for an RF Event. (Message codes prefixed by EVENT_) 
	MessageId byte

	// The code for a specific response or event.
	MessageCode byte
}

func CreateChannelResponseFromMessage(msg *AntCommandMessage) (*ChannelResponse, error) {
	if msg.Id != CHANNEL_RESPONSE_MSG_ID {
		return nil, errors.New("invalid message id")
	}

	return &ChannelResponse{
		AntCommandInfo: newAntCommandInfo(msg.Id, msg.Name),
		ChannelNumber:  msg.Data[0],
		MessageId:      msg.Data[1],
		MessageCode:    msg.Data[2],
	}, nil
}
