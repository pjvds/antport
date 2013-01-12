package messages

import ()

const (
	AssignChannelCommandId   = byte(0x42)
	AssignChannelCommandName = "ASSIGN_CHANNEL"
)

type AssignChannelCommand struct {
	AntCommandInfo

	ChannelNumber byte
	ChannelType   byte
	NetworkNumber byte
}

func CreateAssignChannelCommand(channelNumber byte, channelType byte, networkNumber byte) AssignChannelCommand {
	cmd := newAntCommandInfo(SetChannelIdCommandId, AssignChannelCommandName)
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
