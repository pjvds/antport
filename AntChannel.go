package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntMessage struct {
}

type AntChannel struct {
	device   *usb.Device
	endpoint usb.Endpoint
}

func newAntChannel(usb *usb.Device) *AntChannel {
	epoint, err := usb.OpenEndpoint(1, 0, 0, 1)

	if err != nil {
		log.Fatal(err)
	}

	return &AntChannel{
		device:   usb,
		endpoint: epoint,
	}
}

func (channel *AntChannel) Close() {
	log.Printf("closing ant channel %v", channel.device.Descriptor.Product)

	channel.device.Close()
	channel.device = nil
}

func (channel *AntChannel) ReceiveMessage() *AntMessage {
	return nil
}
