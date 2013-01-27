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
	outEndpoint, err := usbDevice.OpenEndpoint(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_OUT))

	if err != nil {
		return nil, log4go.Error("error opening endpoint: " + err.Error())
	}

	inEndpoint, err := usbDevice.OpenEndpoint(1, 0, 0, uint8(1|usb.ENDPOINT_DIR_IN))

	if err != nil {
		return nil, log4go.Error("error opening endpoint: %s", err)
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

func (device *AntUsbDevice) Reset() {
	log4go.Debug("resetting usb hardware")

SEND_RESET:
	// Hard reset device first
	resetBuffer := []byte{
		0x00, 0x00, 0x00,
		0x00, 0x00, 0x00,
		0x00, 0x00, 0x00,
		0x00, 0x00, 0x00,
		0x00, 0x00, 0x00}

	_, err := device.Write(resetBuffer)
	for err != nil {
		log4go.Warn("error while writing reset bytes to usb device: %v", err.Error())
		_, err = device.Write(resetBuffer)
	}

	buffer := make([]byte, 16)
	// Read hard reset reply
	_, err = device.Read(buffer)
	for err != nil {
		log4go.Warn("error while reading reset bytes reply from usb device: %v", err.Error())
		goto SEND_RESET
	}
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

	log4go.Debug("%v bytes written to usb device", n)
	return n, nil
}

func (channel *AntUsbDevice) Close() {
	if channel.usbDevice != nil {
		log4go.Debug("closing ant channel %v", channel.usbDevice.Descriptor.Product)

		channel.usbDevice.Close()
		channel.usbDevice = nil
	}
}
