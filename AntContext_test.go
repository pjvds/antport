package antport

import (
	"testing"
)

func TestNewContext(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	if ctx == nil {
		t.Error("context was not created")
	}

	if ctx.usb == nil {
		t.Error("context didn create usb context")
	}
}

func TestAntContextClose(t *testing.T) {
	ctx := NewContext()
	ctx.Close()

	if ctx.usb != nil {
		t.Error("context close didn't cleanup usb context")
	}
}

func TestAntContextListChannels(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	channels, err := ctx.ListChannels()

	if err != nil {
		t.Errorf("error while getting channels list: %v", err)
		t.FailNow()
	}

	nChannels := len(channels)

	if nChannels != 1 {
		t.Errorf("found %v channel(s) instead of 1", nChannels)
	}
}
