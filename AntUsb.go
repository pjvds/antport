package antport

import ()

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008

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
