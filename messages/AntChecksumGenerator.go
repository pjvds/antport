package messages

import ()

// XOR of all bytes
func GenerateChecksum(data []byte) byte {
	checksum := byte(0)
	for _, b := range data {
		checksum = checksum ^ b
	}

	return checksum
}
