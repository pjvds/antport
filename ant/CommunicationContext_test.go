package ant

import (
	"github.com/pjvds/antport/hardware"
	"testing"
)

func TestMessageChat(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	ctx := NewCommunicationContext(device)
	defer ctx.Close()
	ctx.Open()

	ctx.Send(RequestMessage(0, MESG_CAPABILITIES_ID)).WaitForReply(func(msg AntMessage) bool {
		return msg.Id == MESG_CAPABILITIES_ID
	})
}

func TestCommunicationContextClose(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	ctx := NewCommunicationContext(device)
	ctx.Open()
	ctx.Close()

	if ctx.device != nil {
		t.Log("Close didn't cleanup device reference.")
		t.Fail()
	}

	if ctx.input != nil {
		t.Log("Close didn't cleanup input channel.")
		t.Fail()
	}

	if ctx.Output != nil {
		t.Log("Close didn't cleanup output channel.")
		t.Fail()
	}
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
