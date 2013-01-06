package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

type AntContext struct {
	usb *usb.Context
}

func NewContext() *AntContext {
	usbCtx := usb.NewContext()

	context := &AntContext{
		usb: usbCtx,
	}

	return context
}

func (ctx *AntContext) ListChannels() ([]*AntChannel, error) {
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

func (ctx *AntContext) Close() {
	//log.Print("closing usb context") <-- CAUSES ERROR :-S

	ctx.usb.Close()
	ctx.usb = nil
}
