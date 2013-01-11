package antport

import (
	"github.com/kylelemons/gousb/usb"
)

type AntUsbEndpoint struct {
	ePoint usb.Endpoint
}

func (endPoint AntUsbEndpoint) Read(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Read(buffer)
}

func (endPoint AntUsbEndpoint) Write(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Write(buffer)
}
