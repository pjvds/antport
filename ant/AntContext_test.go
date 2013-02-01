package ant

import (
	"testing"
)

func TestAntContextInit(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	comm := NewCommunicationContext(device)
	defer comm.Close()

	ctx := NewAntContext(&comm)
	err := ctx.Init()

	if err != nil {
		t.Logf("AntContext Init failed: %v", err.Error())
		t.Fail()
	}
}

func TestAntContextInit_should_create_channels(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	comm := NewCommunicationContext(device)
	defer comm.Close()

	ctx := NewAntContext(&comm)
	err := ctx.Init()

	if err != nil {
		t.Logf("AntContext Init failed: %v", err.Error())
		t.Fail()
	}

	if len(ctx.Channels) == 0 {
		t.Log("AntContext Init didn't create channels")
		t.Fail()
	}
}
