package antport

import(
	"fmt"
	"flag"
	"log"
	
	"github.com/kylelemons/gousb/usb"
)

var (
	debug = flag.Int("debug", 0, "libusb debug level (0..3)")
)

type Context struct {
	usb *usb.Context
}

func Close(ctx *Context) error {
	ctx.usb.Close()
	ctx.usb = nil

	return nil
}

func NewContext() *Context {
	usbCtx := usb.NewContext()
	context := &Context{
		usb: usbCtx,
	}

	return context
}

func Connect(ctx *Context) bool {
	flag.Parse()	

	// Debugging can be turned on; this shows some of the inner workings of the libusb package.
	ctx.usb.Debug(*debug)	

	devs, err := ctx.usb.ListDevices(func(desc *usb.Descriptor) bool {	
		// The usbid package can be used to print out human readable information.
		fmt.Printf("%03d.%03d %s:%s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product)

		// We are looking for the specific vendor and device
		if desc.Vendor == 0xfcf && desc.Product == 0x1008 {
			return true
		}
		
		return false
	})

	// All Devices returned from ListDevices must be closed.
	defer func() {
		for _, d := range devs {
			log.Printf("closing usb context for device %s:%s\n", d.Vendor, d.Product)
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
		return false
	}

	if len(devs) > 1 {
		log.Fatal("multiple devices found")
		return false
	}

	return true
}
