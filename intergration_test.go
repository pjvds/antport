package antport

import (
	"github.com/pjvds/antport/ant"
	"github.com/pjvds/antport/fs"
	"github.com/pjvds/antport/hardware"
	"log"
	"testing"
)

func TestOpenContext(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	ctx := ant.CreateAntContext(device)
	ctx.HardResetSystem(5)
	ctx.Init()
	defer ctx.Close()

	fs := fs.NewAntFsContext(ctx)
	fs.OpenAntsFsSearchChannel()

	reply, err := ctx.ReceiveReply()
	log.Printf("Reply: %s", reply)
	log.Printf("error: %s", err)
}

func GetSingleChannelOrFail(t *testing.T) (*hardware.AntUsbDevice, *hardware.AntUsbContext) {
	ctx := hardware.NewUsbContext()
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

func CloseContextAndFail(t *testing.T, ctx *hardware.AntUsbContext, message string) {
	defer ctx.Close()
	t.Error(message)
	t.FailNow()
}
