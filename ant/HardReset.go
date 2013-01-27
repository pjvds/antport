package ant

import ()

func HardReset() AntMessage {
	return AntMessage{
		// 15 is the buffer we want to send and 4
		// is the number of overhead bytes we want
		// to substract from the data so 15 zero
		// bytes will be send
		Data: make([]byte, 15-4),
	}
}
