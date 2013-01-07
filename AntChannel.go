package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntChannelReader interface {
	Read(buffer []byte) (n int, err error)
}

type AntChannelWriter interface {
	Write(buffer []byte) (n int, err error)
}

type AntChannel interface {
	CreateReader() (reader AntChannelReader, err error)
}

type AntMessage struct {
	raw  []byte
	size int
}

type AntUsbChannel struct {
	device *usb.Device
}

type AntUsbEndpoint struct {
	ePoint usb.Endpoint
}

func newAntChannel(usb *usb.Device) *AntUsbChannel {
	return &AntUsbChannel{
		device: usb,
	}
}

func (channel *AntUsbChannel) CreateReader() (reader AntChannelReader, err error) {
	device := channel.device

	epoint, err := device.OpenEndpoint(1, 0, 0, uint8(1)|uint8(usb.ENDPOINT_DIR_IN))

	if err != nil {
		return nil, err
	}
	return &AntUsbEndpoint{
		ePoint: epoint}, nil
}

func (channel *AntUsbChannel) CreateWriter() (writer AntChannelWriter, err error) {
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

func (channel *AntUsbChannel) Close() {
	log.Printf("closing ant channel %v", channel.device.Descriptor.Product)

	channel.device.Close()
	channel.device = nil
}
