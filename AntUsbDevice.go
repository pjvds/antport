package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

type AntUsbDevice struct {
	device *usb.Device
}

func newAntUsbDevice(usb *usb.Device) *AntUsbDevice {
	return &AntUsbDevice{
		device: usb,
	}
}

func (channel *AntUsbDevice) CreateReader() (reader AntUsbReader, err error) {
	device := channel.device
	log.Printf("opening endpoint for reading")

	epoint, err := device.OpenEndpoint(1, 0, 0, uint8(1)|uint8(usb.ENDPOINT_DIR_IN))

	if err != nil {
		log.Printf("opening failed: " + err.Error())
		return nil, err
	}
	return &AntUsbEndpoint{
		ePoint: epoint}, nil
}

func (channel *AntUsbDevice) CreateWriter() (writer AntUsbWriter, err error) {
	device := channel.device

	log.Printf("opening endpoint for writing")

	epoint, err := device.OpenEndpoint(1, 0, 0, uint8(1)|uint8(usb.ENDPOINT_DIR_OUT))

	if err != nil {
		log.Printf("opening failed: " + err.Error())
		return nil, err
	}
	return &AntUsbEndpoint{
		ePoint: epoint}, nil
}

func (channel *AntUsbDevice) Close() {
	log.Printf("closing ant channel %v", channel.device.Descriptor.Product)

	channel.device.Close()
	channel.device = nil
}
