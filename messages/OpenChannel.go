package messages

import ()

const (
	OPEN_CHANNEL_MSG_ID   = byte(0x4b)
	OpenChannelCommandName = "OPEN_CHANNEL"
)

/* This message is sent to the module to open 
a channel that has been previously assigned and 
configured with the configuration messages. 
Execution of this command causes the channel 
to commence operation, and either data messages 
or events begin to be issued in association with
this channel. */
type OpenChannelCommand struct {
	AntCommandInfo

	ChannelNumber byte
}

func CreateOpenChannelCommand(channelNumber byte) OpenChannelCommand {
	cmd := newAntCommandInfo(SET_CHANNEL_ID_MSG_ID, OpenChannelCommandName)
	return OpenChannelCommand{
		AntCommandInfo: cmd,
		ChannelNumber:  channelNumber,
	}
}

func (cmd OpenChannelCommand) Data() []byte {
	return []byte{
		cmd.ChannelNumber,
	}
}
