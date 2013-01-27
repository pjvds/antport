package ant

import (
	"fmt"
	"github.com/pjvds/antport/hardware"
	"github.com/pjvds/antport/messages"
	"log"
)

// Send AntCommand to AntDevice
type messageSender struct {
	hardware.AntDevice

	maxRetry int
}

func newSender(device hardware.AntDevice) messageSender {
	return messageSender{
		AntDevice: device,
		maxRetry:  25,
	}
}

func (sender messageSender) SendCommand(cmd messages.AntCommand) (ok bool, err error) {
	log.Printf("sending command: %v", cmd.Name())

	msg := messages.NewMessage(cmd)
	data := msg.Pack()
	n, err := sender.Write(data)

	for retries := 1; retries < sender.maxRetry+1; retries++ {
		if err != nil {
			log.Println("error while writing to device: " + err.Error())
			log.Printf("will retry (%v/%v)", retries, sender.maxRetry)

			n, err = sender.Write(data)
		}
	}

	if err != nil {
		log.Println("error while writing to device: " + err.Error())
		return false, err
	}

	if n != len(data) {
		err = fmt.Errorf("number of written bytes (%v) differs from data length (%v)", n, len(data))
		return false, err
	}

	log.Printf("ANT message send: %v", msg.Name)
	return true, nil
}
