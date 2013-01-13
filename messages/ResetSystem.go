package messages

import ()

const (
	RESET_SYSTEM_MSG_ID   = byte(0x4a)
	ResetSystemCommandName = "RESET_SYSTEM"
)

type ResetSystemCommand struct {
	AntCommandInfo
}

func CreateResetSystemCommand() ResetSystemCommand {
	cmd := newAntCommandInfo(SET_CHANNEL_ID_MSG_ID, ResetSystemCommandName)
	return ResetSystemCommand{
		AntCommandInfo: cmd,
	}
}

func (cmd ResetSystemCommand) Data() []byte {
	return make([]byte, 0)
}
