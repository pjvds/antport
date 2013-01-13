package messages

import ()

const (
	SendBroadcastDataCommandId   = byte(0x43)
	SendBroadcastDataCommandName = "SET_CHANNEL_PERIOD"
)

type SendBroadcastDataCommand struct {
	AntCommandInfo

	ChannelNumber byte

	BroadcastData [8]byte
}

func CreateSendBroadcastDataCommand(channelNumber byte, data [8]byte) SendBroadcastDataCommand {
	cmd := newAntCommandInfo(SetChannelIdCommandId, SendBroadcastDataCommandName)
	return SendBroadcastDataCommand{
		AntCommandInfo: cmd,
		BroadcastData:  data,
	}
}

func (cmd SendBroadcastDataCommand) Data() []byte {
	return []byte{
		cmd.ChannelNumber,
		cmd.BroadcastData[0],
		cmd.BroadcastData[1],
		cmd.BroadcastData[2],
		cmd.BroadcastData[3],
		cmd.BroadcastData[4],
		cmd.BroadcastData[5],
		cmd.BroadcastData[6],
		cmd.BroadcastData[7],
	}
}
