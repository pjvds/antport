package messages

import (
	"bytes"
	"encoding/binary"
)

const (
	SetChannelIdCommandId   = byte(0x51)
	SetChannelIdCommandName = "SET_CHANNEL_ID"
)

type SetChannelIdCommand struct {
	AntCommandInfo

	// Data 1
	ChannelNumber byte
	// Data 2 & 3
	DeviceNumber int
	//Data 4
	DeviceTypeId byte
	// Data 5
	TransType byte
}

func CreateSetChannelIdCommand(channelNumber byte, deviceNumber int, deviceTypeId byte, transType byte) SetChannelIdCommand {
	cmd := newAntCommandInfo(SetChannelIdCommandId, SetChannelIdCommandName)
	return SetChannelIdCommand{
		AntCommandInfo: cmd,
		ChannelNumber:  channelNumber,
		DeviceNumber:   deviceNumber,
		DeviceTypeId:   deviceTypeId,
		TransType:      transType,
	}
}

func (cmd SetChannelIdCommand) Data() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, cmd.ChannelNumber)
	binary.Write(buffer, binary.LittleEndian, cmd.DeviceNumber)
	binary.Write(buffer, binary.LittleEndian, cmd.DeviceTypeId)
	binary.Write(buffer, binary.LittleEndian, cmd.TransType)

	return buffer.Bytes()
}
