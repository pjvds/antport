package messages

import (
	"bytes"
	"encoding/binary"
)

const (
	SET_CHANNEL_RF_FREQ_MSG_ID   = byte(0x45)
	SetChannelRfFrequentyCommandName = "SET_CHANNEL_RF_FREQ"
)

type SetChannelRfFrequentyCommand struct {
	AntCommandInfo

	ChannelNumber byte

	ChannelRfFrequenty byte
}

func CreateSetChannelRfFrequentyCommand(channelNumber byte, channelRfFrequenty byte) SetChannelRfFrequentyCommand {
	cmd := newAntCommandInfo(SET_CHANNEL_ID_MSG_ID, SetChannelRfFrequentyCommandName)
	return SetChannelRfFrequentyCommand{
		AntCommandInfo:     cmd,
		ChannelNumber:      channelNumber,
		ChannelRfFrequenty: channelRfFrequenty,
	}
}

func (cmd SetChannelRfFrequentyCommand) Data() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, cmd.ChannelNumber)
	binary.Write(buffer, binary.LittleEndian, cmd.ChannelRfFrequenty)

	return buffer.Bytes()
}
