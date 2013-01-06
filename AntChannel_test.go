package antport

import (
	"testing"
)

func TestAntChannelClose(t *testing.T) {
	channel, ctx := GetSingleChannelOrFail(t)
	defer ctx.Close()

	channel.Close()

	if channel.device != nil {
		t.Error("close doesn't cleanup device")
	}
}

func TestPair(t *testing.T) {
	channel, ctx := GetSingleChannelOrFail(t)
	defer ctx.Close()
	defer channel.Close()

	channel.Pair()
}

func TestCommunication(t *testing.T) {
	channel, ctx := GetSingleChannelOrFail(t)
	defer ctx.Close()
	defer channel.Close()

	channel.SendAck()
	msg := channel.ReceiveMessage()

	if msg == nil {
		t.Error("messages not created")
	}
}

func GetSingleChannelOrFail(t *testing.T) (*AntChannel, *AntContext) {
	ctx := NewContext()

	channels, err := ctx.ListChannels()

	if err != nil {
		t.Errorf("ListChannels failed with error: %v", err)
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

func CloseContextAndFail(t *testing.T, ctx *AntContext, message string) {
	defer ctx.Close()
	t.Error("No channels available")
	t.FailNow()
}
