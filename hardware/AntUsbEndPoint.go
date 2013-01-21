package hardware

import (
	"code.google.com/p/log4go"
	"github.com/pjvds/gousb/usb"
)

type AntUsbEndpoint struct {
	ePoint usb.Endpoint
}

func (endPoint AntUsbEndpoint) Read(buffer []byte) (n int, err error) {
	return endPoint.ePoint.Read(buffer)
}

func (endPoint AntUsbEndpoint) Write(buffer []byte) (n int, err error) {
	log4go.Debug("sending %v to end point", len(buffer))

	n, err = endPoint.ePoint.Write(buffer)

	if err != nil {
		return 0, log4go.Error("error while sending to end point: %s", err)
	}
	log4go.Debug("%v bytes send to end point", n)
	return n, err
}
