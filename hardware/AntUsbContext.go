package hardware

import (
	"code.google.com/p/log4go"
	"flag"
	"github.com/pjvds/gousb/usb"
	"github.com/pjvds/gousb/usbid"
)

var (
	debug = flag.Int("debug", 0, "libusb debug level (0..3)")
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

	// Debugging can be turned on; this shows some of the inner workings of the libusb package.
	usbCtx.Debug(*debug)

	context := &AntUsbContext{
		usb: usbCtx,
	}

	return context
}

func (ctx *AntUsbContext) ListAntUsbDevices() ([]*AntUsbDevice, error) {
	devs, err := ctx.usb.ListDevices(func(desc *usb.Descriptor) bool {
		log4go.Debug("Found %03d.%03d %s:%s %s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product, usbid.Describe(desc))

		// The usbid package can be used to print out human readable information.
		log4go.Debug("  Protocol: %s\n", usbid.Classify(desc))

		// We are looking for the specific vendor and device
		if desc.Vendor == ANT_VENDOR_ID && desc.Product == ANT_PRODUCT_ID {
			log4go.Debug("This is an ANT device")

			// The configurations can be examined from the Descriptor, though they can only
			// be set once the device is opened.  All configuration references must be closed,
			// to free up the memory in libusb.
			for _, cfg := range desc.Configs {
				// This loop just uses more of the built-in and usbid pretty printing to list
				// the USB devices.
				log4go.Debug("  %s:\n", cfg)
				for _, alt := range cfg.Interfaces {
					log4go.Debug("    --------------\n")
					for _, iface := range alt.Setups {
						log4go.Debug("(iface)    %s\n", iface)
						log4go.Debug("(classify)      %s\n", usbid.Classify(iface))
						for _, end := range iface.Endpoints {
							log4go.Debug("(end)      %s\n", end)
							log4go.Debug("	number: %s\n", end.Number())
							log4go.Debug("	address: %s\n", end.Address)
							log4go.Debug("	sync: %s\n", end.SynchAddress)
						}
					}
				}
				log4go.Debug("    --------------\n")
			}

			return true
		}

		return false
	})

	// ListDevices can occaionally fail, so be sure to check its return value.
	if err != nil {
		log4go.Critical("list: %s", err)
		return nil, err
	}

	nDevices := len(devs)
	channels := make([]*AntUsbDevice, nDevices)
	for i := 0; i < len(devs); i++ {
		usbDevice := devs[i]
		channels[i], err = newAntUsbDevice(usbDevice)

		if err != nil {
			panic(err)
		}
	}
	return channels, nil
}

func (ctx *AntUsbContext) Close() {
	// log.Print("closing usb context") // TODO <-- CAUSES ERROR :-S

	ctx.usb.Close()
	ctx.usb = nil
}
