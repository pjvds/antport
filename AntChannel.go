package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
	"time"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntMessage struct {
	raw  []byte
	size int
}

type AntChannel struct {
	device      *usb.Device
	inEndpoint  usb.Endpoint
	outEndpoint usb.Endpoint
}

func newAntChannel(usb *usb.Device) *AntChannel {
	var conf, iface, setup, epoint uint8 = 1, 0, 0, 1

	log.Println("opening out endpoint")
	outEpoint, err := usb.OpenEndpoint(conf, iface, setup, epoint)
	if err != nil {
		log.Fatal("error while opening out endpoint: " + err.Error())
	}

	log.Println("end point created")

	return &AntChannel{
		device:      usb,
		outEndpoint: outEpoint,
		inEndpoint:  nil,
	}
}

func (channel *AntChannel) Pair() {
	inEndpointOpen := false

	for !inEndpointOpen {
		channel.SendAck()
		log.Println("opening in endpoint")
		inEpoint, err := channel.device.OpenEndpoint(1, 0, 0, 0x81)
		if err != nil {
			log.Println("error while opening in endpoint: " + err.Error())

			time.Sleep(500)
		} else {
			inEndpointOpen = true
			channel.inEndpoint = inEpoint
		}
	}
}

func (channel *AntChannel) Close() {
	log.Printf("closing ant channel %v", channel.device.Descriptor.Product)

	channel.device.Close()
	channel.device = nil
}

func (channel *AntChannel) SendAck() {
	buffer := []byte{0x44, 0x02, 0x07, 0x04, 0x00, 0x00, 0x00, 0x00}
	channel.write(buffer)
}

func (channel *AntChannel) write(content []byte) {
	log.Println("starting to send")
	bytesWritten, err := channel.outEndpoint.Write(content)

	if bytesWritten == 0 {
		log.Fatal("nothing send!")
	} else {
		log.Printf("%v bytes send\n", bytesWritten)
	}

	if err != nil {
		log.Fatal("error while sending to channel endpoint: " + err.Error())
	}
}

func (channel *AntChannel) ReceiveMessage() *AntMessage {
	var buffer = make([]byte, 255)

	bytesRead, err := channel.inEndpoint.Read(buffer)

	if err != nil {
		log.Fatal("error while reading from channel endpoint: " + err.Error())
	}

	return &AntMessage{
		raw:  buffer,
		size: bytesRead,
	}
}
