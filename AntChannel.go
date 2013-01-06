package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntChannel struct {
	device *usb.Device
}

func (channel *AntChannel) Close() {
	log.Printf("closing ant channel %v", channel.device.Descriptor.Product)

	channel.device.Close()
	channel.device = nil
}
