package antport

import (
	"testing"
)

func TestAntChannelClose(t *testing.T) {
	channel, err := GetSingleChannelOrFail(t)
	channel.Close()

	if err == nil {
		t.Errorf("close failed with error: %v", err)
	}

	if channel.device != nil {
		t.Error("close doesn't cleanup device")
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
