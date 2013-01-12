package antport

type AntCommand struct {
	SYNC      byte
	Direction string
	Id        byte
	Name      string
	Data      []byte
}

func (command *AntCommand) Pack() []byte {
	overheadSize := MESG_SYNC_SIZE + MESG_SIZE_SIZE + MESG_ID_SIZE + MESG_CHECKSUM_SIZE
	dataSize := byte(len(command.Data))
	dataOffset := byte(3)
	data := make([]byte, overheadSize+dataSize)

	// Set message values
	data[0] = command.SYNC
	data[1] = dataSize
	data[2] = command.Id

	// Set message data/payload
	for i := byte(0); i < dataSize; i++ {
		data[i+dataOffset] = command.Data[i]
	}

	// Set checksum at last byte
	data[len(data)-1] = GenerateChecksum(data)
	return data
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

func CreateSetChannelIdCommand(channelNumber byte, deviceNumber byte, deviceTypeId byte, transType byte) {
	data := []byte{
		channelNumber,
		deviceNumber,
		deviceTypeId,
		transType,
	}
	return newMessage(DIR_OUT, 0x51, "SET_CHANNEL_ID", data)
}
