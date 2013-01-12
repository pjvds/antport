package messages

import ()

const (
	RequestMessageCommandId = byte(0x4d)
)

type RequestMessageCommand struct {
	AntCommandInfo
	Channel   byte
	MessageId byte
}

func CreateRequestMessageCommand(channel byte, messageId byte) RequestMessageCommand {
	cmd := newAntCommandInfo(RequestMessageCommandId, "REQUEST_MESSAGE")
	return RequestMessageCommand{
		AntCommandInfo: cmd,
		Channel:    channel,
		MessageId:  messageId,
	}
}

func (cmd RequestMessageCommand) data() []byte {
	return []byte{
		cmd.Channel,
		cmd.MessageId,
	}
}
