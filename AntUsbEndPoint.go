package antport

import (
	"github.com/kylelemons/gousb/usb"
	"log"
)

type AntUsbEndpoint struct {
	ePoint usb.Endpoint
}

func (endPoint AntUsbEndpoint) Read(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Read(buffer)
}

func (endPoint AntUsbEndpoint) Write(buffer []byte) (n int, err error) {
	log.Printf("sending %v to end point", len(buffer))

	n, err = endPoint.ePoint.Write(buffer)

	if err != nil {
		log.Println("error while sending to end point: " + err.Error())
	} else {
		log.Printf("%v bytes send to end point", n)
	}

	return n, err
}
