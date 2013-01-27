package ant

import ()

type AntMessage struct {
	Sync   byte
	Id     byte
	Data   []byte
	Length byte
}

func NewAntMessage(sync byte, id byte, data []byte) AntMessage {
	return AntMessage{
		Sync:   sync,
		Id:     id,
		Data:   data,
		Length: byte(len(data)),
	}
}

// XOR of all bytes
func GenerateChecksum(data []byte) byte {
	checksum := byte(0)
	for _, b := range data {
		checksum = checksum ^ b
	}

	return checksum
}

func (msg AntMessage) ToBytes() []byte {
	payload := []byte{msg.Sync, msg.Length, msg.Id}
	payload = append(payload, msg.Data...)

	chechsum := GenerateChecksum(payload)
	payload = append(payload, chechsum)

	return payload
}
