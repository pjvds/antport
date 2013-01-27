package ant

import (
	"fmt"
	"github.com/pjvds/antport/hardware"
	"log"
)

// Send AntCommand to AntDevice
type MessageSender struct {
	hardware.AntDevice

	maxRetry int
}

func newSender(device hardware.AntDevice) MessageSender {
	return MessageSender{
		AntDevice: device,
		maxRetry:  25,
	}
}

func (sender MessageSender) Send(msg AntMessage) (err error) {
	log.Printf("sending message: %v", msg.Id)

	data := msg.ToBytes()
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
		return err
	}

	if n != len(data) {
		err = fmt.Errorf("number of written bytes (%v) differs from data length (%v)", n, len(data))
		return err
	}

	log.Printf("ANT message send: %v", msg.Id)
	return nil
}
