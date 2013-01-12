package messages

import ()

const (
	RequestMessageCommandId   = byte(0x4d)
	RequestMessageCommandName = "REQUEST_MESSAGE"
)

type RequestMessageCommand struct {
	AntCommandInfo
	Channel   byte
	MessageId byte
}

func CreateRequestMessageCommand(channel byte, messageId byte) RequestMessageCommand {
	cmd := newAntCommandInfo(RequestMessageCommandId, RequestMessageCommandName)
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
