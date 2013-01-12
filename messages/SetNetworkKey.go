package messages

import (
	"bytes"
	"encoding/binary"
)

const (
	SetNetworkKeyCommandId   = byte(0x46)
	SetNetworkKeyCommandName = "SET_NETWORK_KEY"
)

type SetNetworkKeyCommand struct {
	AntCommandInfo

	NetworkNumber byte
	NetworkKey    [8]byte
}

func CreateSetNetworkKeyCommand(networkNumber byte, networkKey [8]byte) SetNetworkKeyCommand {
	cmd := newAntCommandInfo(SetNetworkKeyCommandId, SetNetworkKeyCommandName)
	return SetNetworkKeyCommand{
		AntCommandInfo: cmd,
		NetworkNumber:  networkNumber,
		NetworkKey:     networkKey,
	}
}

func (cmd SetNetworkKeyCommand) Data() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, cmd.NetworkNumber)
	binary.Write(buffer, binary.LittleEndian, cmd.NetworkKey)

	return buffer.Bytes()
}
