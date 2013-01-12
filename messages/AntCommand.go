package messages

import ()

type AntCommandInfo struct {
	id   byte
	name string
}

type AntCommand interface {
	Id() byte
	Name() string
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
		id:   id,
		name: name,
	}
}

func (cmd AntCommandInfo) Id() byte {
	return cmd.id
}

func (cmd AntCommandInfo) Name() string {
	return cmd.name
}
