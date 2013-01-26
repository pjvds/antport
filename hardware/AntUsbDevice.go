package hardware

import (
	"code.google.com/p/log4go"
	"github.com/pjvds/gousb/usb"
	"time"
)

type AntUsbDevice struct {
	usbDevice   *usb.Device
	inEndpoint  usb.Endpoint
	outEndpoint usb.Endpoint
}

func newAntUsbDevice(usbDevice *usb.Device) (*AntUsbDevice, error) {
	log4go.Debug("creating new AntUsbDevice")

	log4go.Debug("opening IN endpoint...")
	inEndpoint, err := usbDevice.OpenEndpoint(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_IN))

	if err != nil {
		return nil, log4go.Error("error opening endpoint: %s", err)
	} else {
		log4go.Debug("IN endpoint opened succesfully")
	}

	log4go.Debug("opening OUT endpoint...")
	outEndpoint, err := usbDevice.OpenEndpoint(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_OUT))

	if err != nil {
		return nil, log4go.Error("error opening endpoint: " + err.Error())
	} else {
		log4go.Debug("IN endpoint opened succesfully")
	}

	usbDevice.WriteTimeout = time.Second * 2
	usbDevice.ReadTimeout = time.Second * 2

	log4go.Debug("AntUsbDevice created succesfully")
	return &AntUsbDevice{
		usbDevice:   usbDevice,
		inEndpoint:  inEndpoint,
		outEndpoint: outEndpoint,
	}, nil
}

func (device *AntUsbDevice) Read(buffer []byte) (int, error) {
	epoint := device.inEndpoint
	n, err := epoint.Read(buffer)

	if err != nil {
		return 0, log4go.Error("error while reading from device: %s", err)
	}

	return n, nil
}

func (device *AntUsbDevice) Write(data []byte) (int, error) {
	epoint := device.outEndpoint
	n, err := epoint.Write(data)

	if err != nil {
		return 0, log4go.Error("error while writing to device: %s", err)
	}

	return n, nil
}

func (channel *AntUsbDevice) Close() {
	log4go.Debug("closing ant channel %v", channel.usbDevice.Descriptor.Product)

	channel.usbDevice.Close()
	channel.usbDevice = nil
}
