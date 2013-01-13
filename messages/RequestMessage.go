package messages

import ()

const (
	REQUEST_MESSAGE_MSG_ID   = byte(0x4d)
	RequestMessageCommandName = "REQUEST_MESSAGE"
)

type RequestMessageCommand struct {
	AntCommandInfo
	Channel   byte
	MessageId byte
}

func CreateRequestMessageCommand(channel byte, messageId byte) RequestMessageCommand {
	cmd := newAntCommandInfo(REQUEST_MESSAGE_MSG_ID, RequestMessageCommandName)
	return RequestMessageCommand{
		AntCommandInfo: cmd,
		Channel:        channel,
		MessageId:      messageId,
	}
}

func (cmd RequestMessageCommand) Data() []byte {
	return []byte{
		cmd.Channel,
		cmd.MessageId,
	}
}
