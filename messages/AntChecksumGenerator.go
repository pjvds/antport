package messages

import ()

type AntChecksumGenerator interface {
	GenerateChecksum(data []byte) byte
}

func GenerateChecksum(data []byte) byte {
	checksum := byte(0)
	for i := 0; i < len(data); i++ {
		checksum = checksum ^ data[i]
	}

	return checksum
}
