package antport

import(
	"fmt"
	"flag"
	"log"
	
	"github.com/kylelemons/gousb/usb"
	"github.com/kylelemons/gousb/usbid"
)

var (
	debug = flag.Int("debug", 0, "libusb debug level (0..3)")
)

func IsHardwarePresent() bool {
	flag.Parse()

	// Only one context should be needed for an application.  It should always be closed.
	ctx := usb.NewContext()
	defer ctx.Close()

	// Debugging can be turned on; this shows some of the inner workings of the libusb package.
	ctx.Debug(*debug)	

	devs, err := ctx.ListDevices(func(desc *usb.Descriptor) bool {	
		// The usbid package can be used to print out human readable information.
		fmt.Printf("%03d.%03d %s:%s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product)

		if desc.Vendor == 0xfcf {
		fmt.Printf("  Protocol: %s\n", usbid.Classify(desc))
		// The configurations can be examined from the Descriptor, though they can only
		// be set once the device is opened.  All configuration references must be closed,
		// to free up the memory in libusb.
		for _, cfg := range desc.Configs {
			// This loop just uses more of the built-in and usbid pretty printing to list
			// the USB devices.
			fmt.Printf("  %s:\n", cfg)
			for _, alt := range cfg.Interfaces {
				fmt.Printf("    --------------\n")
				for _, iface := range alt.Setups {
					fmt.Printf("    %s\n", iface)
					fmt.Printf("      %s\n", usbid.Classify(iface))
					for _, end := range iface.Endpoints {
						fmt.Printf("      %s\n", end)
					}
				}
			}
			fmt.Printf("    --------------\n")
			}
			return true
		}
		
		return false
	})

	// All Devices returned from ListDevices must be closed.
	defer func() {
		for _, d := range devs {
			log.Printf("closing device %s", d)
			d.Close()
		}
	}()

	// ListDevices can occaionally fail, so be sure to check its return value.
	if err != nil {
		log.Fatalf("list: %s", err)
		return false
	}

	if len(devs) == 0 {
		log.Fatal("no devices found")
	}

	for _, d := range devs {
		log.Printf("Working with devices %s", d)
	}
		
	numberOfDevices := len(devs)
	log.Printf("%v device(s) found.", numberOfDevices)
	return numberOfDevices > 0
}
