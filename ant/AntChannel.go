package ant

import ()

type AntChannel struct {
	Number byte
}

func NewAntChannel(number byte) AntChannel {
	return AntChannel{
		Number: number,
	}
}
