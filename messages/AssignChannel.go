package messages

import ()

const (
	ASSIGN_CHANNEL_MSG_ID   = byte(0x42)
	ASSIGN_CHANNEL_MSG_NAME = "ASSIGN_CHANNEL"
)

type AssignChannelCommand struct {
	AntCommandInfo

	ChannelNumber byte
	ChannelType   byte
	NetworkNumber byte
}

func CreateAssignChannelCommand(channelNumber byte, channelType byte, networkNumber byte) AssignChannelCommand {
	cmd := newAntCommandInfo(SET_CHANNEL_ID_MSG_ID, ASSIGN_CHANNEL_MSG_NAME)
	return AssignChannelCommand{
		AntCommandInfo: cmd,
		ChannelNumber:  channelNumber,
		ChannelType:    channelType,
		NetworkNumber:  networkNumber,
	}
}

func (cmd AssignChannelCommand) Data() []byte {
	return []byte{
		cmd.ChannelNumber,
		cmd.ChannelType,
		cmd.NetworkNumber,
	}
}
