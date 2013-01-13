package hardware

import (
	"github.com/kylelemons/gousb/usb"
	"github.com/kylelemons/gousb/usbid"
	"log"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntUsbContext struct {
	usb *usb.Context
}

func NewUsbContext() *AntUsbContext {
	usbCtx := usb.NewContext()

	context := &AntUsbContext{
		usb: usbCtx,
	}

	return context
}

func (ctx *AntUsbContext) ListAntUsbDevices() ([]*AntUsbDevice, error) {
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
						log.Printf("(iface)    %s\n", iface)
						log.Printf("(classify)      %s\n", usbid.Classify(iface))
						for _, end := range iface.Endpoints {
							log.Printf("(end)      %s\n", end)
							log.Printf("	number: %s\n", end.Number())
							log.Printf("	address: %s\n", end.Address)
							log.Printf("	sync: %s\n", end.SynchAddress)
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
	channels := make([]*AntUsbDevice, nDevices)
	for i := 0; i < len(devs); i++ {
		usbDevice := devs[i]

		channels[i] = newAntUsbDevice(usbDevice)
	}
	return channels, nil
}

func (ctx *AntUsbContext) Close() {
	// log.Print("closing usb context") // TODO <-- CAUSES ERROR :-S

	ctx.usb.Close()
	ctx.usb = nil
}
