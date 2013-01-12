package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

type AntUsbDevice struct {
	usbDevice *usb.Device
}

type AntDevice interface {
	Read(buffer []byte) (n int, err error)
	Writer(data []byte) (n int, err error)
}

func newAntUsbDevice(usbDevice *usb.Device) *AntUsbDevice {
	return &AntUsbDevice{
		usbDevice: usbDevice,
	}
}

func (device *AntUsbDevice) Read(buffer []byte) (int, error) {
	ep, err := device.usbDevice.OpenEndpoint(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_IN))

	if err != nil {
		log.Println("error while opening end point for reading: " + err.Error())
		return 0, err
	}

	n, err := ep.Read(buffer)

	if err != nil {
		log.Println("error while reading from device: " + err.Error())
	} else {
		log.Printf("%v bytes read from device", n)
	}

	return n, err
}

func (device *AntUsbDevice) Write(data []byte) (int, error) {
	log.Printf("writing %v bytes to device", len(data))

	log.Println("opening end point for writing")
	ep, err := device.usbDevice.OpenEndpoint(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_OUT))

	if err != nil {
		log.Println("error while opening end point for write: " + err.Error())
		return 0, err
	}

	n, err := ep.Write(data)

	if err != nil {
		log.Println("error while writing to device: " + err.Error())
	} else {
		log.Printf("%v bytes written to device", n)
	}

	return n, err
}

func (channel *AntUsbDevice) Close() {
	log.Printf("closing ant channel %v", channel.usbDevice.Descriptor.Product)

	channel.usbDevice.Close()
	channel.usbDevice = nil
}
