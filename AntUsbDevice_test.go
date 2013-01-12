package antport

import (
	"testing"
)

func TestOpenContext(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	resetCommand := CreateResetCommand()
	device.Write(resetCommand.Pack())
	device.Read(make([]byte, 8))
}
