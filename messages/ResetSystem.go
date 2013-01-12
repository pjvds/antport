package messages

import ()

const (
	ResetSystemCommandId   = byte(0x4a)
	ResetSystemCommandName = "RESET_SYSTEM"
)

type ResetSystemCommand struct {
	AntCommandInfo
}

func CreateResetSystemCommand() ResetSystemCommand {
	cmd := newAntCommandInfo(SetChannelIdCommandId, ResetSystemCommandName)
	return ResetSystemCommand{
		AntCommandInfo: cmd,
	}
}

func (cmd ResetSystemCommand) Data() []byte {
	return make([]byte, 0)
}
