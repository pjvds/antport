package messages

import (
	"bytes"
	"encoding/binary"
)

const (
	SET_CHANNEL_SEARCH_WAVEFORM_MSG_ID  = byte(0x49)
	SetChannelSearchWaveformCommandName = "SET_CHANNEL_SEARCH_WAVEFORM"
)

type SetChannelSearchWaveformCommand struct {
	AntCommandInfo

	ChannelNumber byte

	ChannelSearchWaveform uint16
}

func CreateSetChannelSearchWaveformCommand(channelNumber byte, channelSearchWaveform uint16) SetChannelSearchWaveformCommand {
	cmd := newAntCommandInfo(SET_CHANNEL_ID_MSG_ID, SetChannelSearchWaveformCommandName)
	return SetChannelSearchWaveformCommand{
		AntCommandInfo:        cmd,
		ChannelNumber:         channelNumber,
		ChannelSearchWaveform: channelSearchWaveform,
	}
}

func (cmd SetChannelSearchWaveformCommand) Data() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, cmd.ChannelNumber)
	binary.Write(buffer, binary.LittleEndian, cmd.ChannelSearchWaveform)

	return buffer.Bytes()
}
