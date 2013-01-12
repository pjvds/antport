package antport

import (
	"testing"
)

func TestOpenContext(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	ctx := CreateAntContext(device)
	ctx.Init()
	ctx.HardResetSystem()
}

func GetSingleChannelOrFail(t *testing.T) (*AntUsbDevice, *AntUsbContext) {
	ctx := NewUsbContext()
	channels, err := ctx.ListAntUsbDevices()

	if err != nil {
		CloseContextAndFail(t, ctx, "ListChannels failed with error: %v")

	}

	if len(channels) == 0 {
		CloseContextAndFail(t, ctx, "no channels available")
	}

	if len(channels) > 1 {
		CloseContextAndFail(t, ctx, "multiple channels available")
	}

	channel := channels[0]
	return channel, ctx
}

func CloseContextAndFail(t *testing.T, ctx *AntUsbContext, message string) {
	defer ctx.Close()
	t.Error(message)
	t.FailNow()
}
