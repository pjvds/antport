package messages

import ()

const (
	SEND_BROADCAST_DATA_MSG_ID   = byte(0x4e)
	SendBroadcastDataCommandName = "SEND_BROADCAST_DATA"
)

type SendBroadcastDataCommand struct {
	AntCommandInfo

	ChannelNumber byte

	BroadcastData [8]byte
}

func CreateSendBroadcastDataCommand(channelNumber byte, data [8]byte) SendBroadcastDataCommand {
	cmd := newAntCommandInfo(SET_CHANNEL_ID_MSG_ID, SendBroadcastDataCommandName)
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
