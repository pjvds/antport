package messages

import (
	"encoding/hex"
)

var (
	KnownMessageNames = map[byte]string{
		RequestMessageCommandId:   RequestMessageCommandName,
		ResetSystemCommandId:      ResetSystemCommandName,
		SetChannelIdCommandId:     SetChannelIdCommandName,
		CapabilitiesCommandId:     CapabilitiesCommandName,
		SetChannelPeriodCommandId: SetChannelPeriodCommandName,
	}
)

func CommandIdToName(id byte) string {
	name, ok := KnownMessageNames[id]

	if ok {
		return name
	}

	input := []byte{id}
	return "UNKNOWN_" + hex.EncodeToString(input)
}
