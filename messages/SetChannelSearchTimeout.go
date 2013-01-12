package messages

import ()

const (
	SetChannelSearchTimeoutCommandId   = byte(0x44)
	SetChannelSearchTimeoutCommandName = "SET_CHANNEL_SEARCH_TIMEOUT"
)

type SetChannelSearchTimeoutCommand struct {
	AntCommandInfo

	ChannelNumber byte

	/* The search timeout to be used with by this channel for 
	receive searching. A value of 0 will result is no timeout. 
	Each count in this parameter is equivalent to 2.5 seconds.  
	I.E. 240 = 600 seconds = 10 minutes

	The default is 12 (30 seconds) */
	SearchTimeout byte
}

func CreateSetChannelSearchTimeoutCommand(channelNumber byte, searchTimeout byte) SetChannelSearchTimeoutCommand {
	cmd := newAntCommandInfo(SetChannelIdCommandId, SetChannelSearchTimeoutCommandName)
	return SetChannelSearchTimeoutCommand{
		AntCommandInfo: cmd,
		ChannelNumber:  channelNumber,
		SearchTimeout:  searchTimeout,
	}
}

func (cmd SetChannelSearchTimeoutCommand) Data() []byte {
	return []byte{
		cmd.ChannelNumber,
		cmd.SearchTimeout,
	}
}
