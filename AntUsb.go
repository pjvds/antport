package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntUsbReader interface {
	Read(buffer []byte) (n int, err error)
}

type AntUsbWriter interface {
	Write(buffer []byte) (n int, err error)
}

type AntChannel interface {
	CreateReader() (reader AntUsbReader, err error)
}

type AntMessage struct {
	raw  []byte
	size int
}

type AntUsbDevice struct {
	device *usb.Device
}

type AntUsbEndpoint struct {
	ePoint usb.Endpoint
}

func newAntUsbDevice(usb *usb.Device) *AntUsbDevice {
	return &AntUsbDevice{
		device: usb,
	}
}

func (channel *AntUsbDevice) CreateReader() (reader AntUsbReader, err error) {
	device := channel.device

	epoint, err := device.OpenEndpoint(1, 0, 0, uint8(1)|uint8(usb.ENDPOINT_DIR_IN))

	if err != nil {
		return nil, err
	}
	return &AntUsbEndpoint{
		ePoint: epoint}, nil
}

func (channel *AntUsbDevice) CreateWriter() (writer AntUsbWriter, err error) {
	device := channel.device

	epoint, err := device.OpenEndpoint(1, 0, 0, uint8(1)|uint8(usb.ENDPOINT_DIR_OUT))

	if err != nil {
		return nil, err
	}
	return &AntUsbEndpoint{
		ePoint: epoint}, nil
}

func (endPoint AntUsbEndpoint) Write(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Write(buffer)
}

func (endPoint AntUsbEndpoint) Read(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Read(buffer)
}

func (channel *AntUsbDevice) Close() {
	log.Printf("closing ant channel %v", channel.device.Descriptor.Product)

	channel.device.Close()
	channel.device = nil
}
