package antport

import (
	"github.com/kylelemons/gousb/usb"
)

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

type AntUsbEndpoint struct {
	ePoint usb.Endpoint
}

type AntCommand struct {
	SYNC byte
	ID   byte
	NAME string
	DATA []byte
}

func ResetCommand() *AntCommand {
	return newAntMessage(MESG_SYSTEM_RESET_ID, "RESET_SYSTEM")
}

func (command *AntCommand) Pack() []byte {
	overheadSize := MESG_SYNC_SIZE + MESG_SIZE_SIZE + MESG_ID_SIZE + MESG_CHECKSUM_SIZE
	dataSize := byte(len(command.DATA))
	dataOffset := byte(3)
	data := make([]byte, overheadSize+dataSize)

	// Set message values
	data[0] = command.SYNC
	data[1] = dataSize
	data[2] = command.ID

	// Set message data/payload
	for i := byte(0); i < dataSize; i++ {
		data[i+dataOffset] = command.DATA[i]
	}

	// Set checksum
	data[overheadSize+dataSize] = GenerateChecksum(data)
	return data
}

func GenerateChecksum(data []byte) byte {
	checksum := byte(0)
	for i := 0; i < len(data)-1; i++ {
		checksum = checksum ^ data[i]
	}

	return checksum
}

func CreateCapabilitiesCommand() *AntCommand {
	return newAntMessage(0x54, "CAPABILITIES")
}

func CreateResetCommand() *AntCommand {
	return newAntMessage(MESG_SYSTEM_RESET_ID, "SYSTEM_RESET")
}

func SendCommand(writer AntUsbWriter, command *AntCommand) {
	data := command.Pack()
	writer.Write(data)
}

func newAntMessage(id byte, name string) *AntCommand {
	return &AntCommand{
		ID:   id,
		NAME: name,
	}
}

func (endPoint AntUsbEndpoint) Write(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Write(buffer)
}

func (endPoint AntUsbEndpoint) Read(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Read(buffer)
}
