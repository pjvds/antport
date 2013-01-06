package antport

import (
	"log"

	"github.com/kylelemons/gousb/usb"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntChannel struct {
	device *usb.Device
}

func ListChannels(ctx *AntPortContext) ([]*AntChannel, error) {
	devs, err := ctx.usb.ListDevices(func(desc *usb.Descriptor) bool {
		// The usbid package can be used to print out human readable information.
		log.Printf("%03d.%03d %s:%s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product)

		// We are looking for the specific vendor and device
		if desc.Vendor == ANT_VENDOR_ID && desc.Product == ANT_PRODUCT_ID {
			return true
		}

		return false
	})

	// ListDevices can occaionally fail, so be sure to check its return value.
	if err != nil {
		log.Fatalf("list: %s", err)
		return nil, err
	}

	nDevices := len(devs)
	channels := make([]*AntChannel, nDevices)
	for i := 0; i < len(devs); i++ {
		usbDevice := devs[i]
		chnl := &AntChannel{
			device: usbDevice,
		}

		channels[i] = chnl
	}
	return channels, nil
}
