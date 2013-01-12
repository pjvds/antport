package antport

type AntCommand struct {
	SYNC      byte
	Direction string
	Id        byte
	Name      string
	Data      []byte
}

func CreateResetCommand() *AntCommand {
	return newMessage(DIR_OUT, MESG_SYSTEM_RESET_ID, "RESET_SYSTEM", make([]byte, 0))
}

func CreateRequestMessageCommand(channel byte, messageId byte) *AntCommand {
	data := []byte{
		channel, messageId,
	}
	return newMessage(DIR_OUT, 0x4d, "REQUEST_MESSAGE", data)
}
