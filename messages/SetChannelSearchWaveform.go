package messages

import (
	"bytes"
	"encoding/binary"
)

const (
	SetChannelSearchWaveformCommandId   = byte(0x49)
	SetChannelSearchWaveformCommandName = "SET_CHANNEL_SEARCH_WAVEFORM"
)

type SetChannelSearchWaveformCommand struct {
	AntCommandInfo

	ChannelNumber byte

	ChannelSearchWaveform uint16
}

func CreateSetChannelSearchWaveformCommand(channelNumber byte, channelSearchWaveform uint16) SetChannelSearchWaveformCommand {
	cmd := newAntCommandInfo(SetChannelIdCommandId, SetChannelSearchWaveformCommandName)
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
