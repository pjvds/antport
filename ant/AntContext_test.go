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
