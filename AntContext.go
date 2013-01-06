package antport

import (
	"github.com/kylelemons/gousb/usb"
	"github.com/kylelemons/gousb/usbid"
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
		log.Printf("Found %03d.%03d %s:%s %s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product, usbid.Describe(desc))

		// The usbid package can be used to print out human readable information.
		log.Printf("  Protocol: %s\n", usbid.Classify(desc))

		// We are looking for the specific vendor and device
		if desc.Vendor == ANT_VENDOR_ID && desc.Product == ANT_PRODUCT_ID {
			log.Println("This is an ANT device")

			// The configurations can be examined from the Descriptor, though they can only
			// be set once the device is opened.  All configuration references must be closed,
			// to free up the memory in libusb.
			for _, cfg := range desc.Configs {
				// This loop just uses more of the built-in and usbid pretty printing to list
				// the USB devices.
				log.Printf("  %s:\n", cfg)
				for _, alt := range cfg.Interfaces {
					log.Printf("    --------------\n")
					for _, iface := range alt.Setups {
						log.Printf("    %s\n", iface)
						log.Printf("      %s\n", usbid.Classify(iface))
						for _, end := range iface.Endpoints {
							log.Printf("      %s\n", end)
						}
					}
				}
				log.Printf("    --------------\n")
			}

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

		channels[i] = newAntChannel(usbDevice)
	}
	return channels, nil
}

func (ctx *AntContext) Close() {
	// log.Print("closing usb context") // TODO <-- CAUSES ERROR :-S

	ctx.usb.Close()
	ctx.usb = nil
}
