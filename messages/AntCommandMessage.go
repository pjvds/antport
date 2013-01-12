package messages

const (
	DIR_IN  = "IN"
	DIR_OUT = "OUT"
)

type AntCommandMessage struct {
	SYNC      byte
	Direction string
	Id        byte
	Name      string
	Data      []byte
}

func NewMessage(cmd AntCommand) AntCommandMessage {
	id := cmd.Id()
	name := cmd.Name()
	data := cmd.Data()

	return AntCommandMessage{
		SYNC:      MESG_TX_SYNC,
		Direction: DIR_OUT,
		Id:        id,
		Name:      name,
		Data:      data,
	}
}

func (message *AntCommandMessage) Pack() []byte {
	overheadSize := MESG_SYNC_SIZE + MESG_SIZE_SIZE + MESG_ID_SIZE + MESG_CHECKSUM_SIZE
	dataSize := byte(len(message.Data))
	dataOffset := byte(3)
	data := make([]byte, overheadSize+dataSize)

	// Set message values
	data[0] = message.SYNC
	data[1] = dataSize
	data[2] = message.Id

	// Set message data/payload
	for i := byte(0); i < dataSize; i++ {
		data[i+dataOffset] = message.Data[i]
	}

	// Set checksum at last byte
	data[len(data)-1] = GenerateChecksum(data)
	return data
}
