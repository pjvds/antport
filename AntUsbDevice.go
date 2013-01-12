package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

type AntUsbDevice struct {
	usbDevice   *usb.Device
	inEndpoint  usb.Endpoint
	outEndpoint usb.Endpoint
}

type AntDevice interface {
	Read(buffer []byte) (n int, err error)
	Writer(data []byte) (n int, err error)
}

func newAntUsbDevice(usbDevice *usb.Device) *AntUsbDevice {
	inEndpoint, err := usbDevice.OpenEndpoint(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_IN))

	if err != nil {
		log.Println("error opening endpoint: " + err.Error())
	}

	outEndpoint, err := usbDevice.OpenEndpointNoCheck(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_OUT))

	if err != nil {
		log.Println("error opening endpoint: " + err.Error())
	}

	return &AntUsbDevice{
		usbDevice:   usbDevice,
		inEndpoint:  inEndpoint,
		outEndpoint: outEndpoint,
	}
}

func (device *AntUsbDevice) Read(buffer []byte) (int, error) {
	epoint := device.inEndpoint
	n, err := epoint.Read(buffer)

	if err != nil {
		log.Println("error while reading from device: " + err.Error())
	} else {
		log.Printf("%v bytes read from device", n)
	}

	return n, err
}

func (device *AntUsbDevice) Write(data []byte) (int, error) {
	epoint := device.outEndpoint
	n, err := epoint.Write(data)

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
