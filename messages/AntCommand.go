package messages

import ()

type AntCommandInfo struct {
	Id   byte
	Name string
}

type AntCommand interface {
	Data() []byte
}

const (
	MESG_TX_SYNC          = byte(0xA4)
	MESG_RX_SYNC          = byte(0xA5)
	MESG_SYNC_SIZE        = byte(1)
	MESG_SIZE_SIZE        = byte(1)
	MESG_ID_SIZE          = byte(1)
	MESG_CHANNEL_NUM_SIZE = byte(1)
	MESG_EXT_MESG_BF_SIZE = byte(1) // NOTE: this could increase in the future
	MESG_CHECKSUM_SIZE    = byte(1)
	MESG_DATA_SIZE        = byte(9)

	MESG_SYSTEM_RESET_ID   = byte(0x4A)
	MESG_SYSTEM_RESET_SIZE = byte(1)
)

// Creates a new AntCommandInfo message
func newAntCommandInfo(id byte, name string) AntCommandInfo {
	return AntCommandInfo{
		Id:   id,
		Name: name,
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
