package ant

import (
	"github.com/pjvds/antport/hardware"
	"github.com/pjvds/antport/messages"
	"log"
)

type messageReceiver struct {
	hardware.AntDevice

	maxRetry int
}

func newReceiver(device hardware.AntDevice) messageReceiver {
	return messageReceiver{
		AntDevice: device,
		maxRetry:  25,
	}
}

func (receiver messageReceiver) ReceiveReply() (reply *messages.AntCommandMessage, err error) {
	log.Println("receiving reply...")

	buffer := make([]byte, 16)
	n, err := receiver.Read(buffer)

	for retries := 1; retries < receiver.maxRetry+1; retries++ {
		if err != nil {
			log.Printf("error while receiving reply, %v bytes read: %s", n, err)
			log.Printf("will retry (%v/%v)", retries, receiver.maxRetry)

			n, err = receiver.Read(buffer)
		}
	}

	if err != nil {
		log.Println("error reading from device: " + err.Error())
		return nil, err
	}

	data := make([]byte, 0)
	name := messages.InMessageIdToName(buffer[2])
	size := buffer[1]

	log.Printf("ANT message received: %v", name)
	if size > 0 {
		data = buffer[3:size]
	}

	reply = &messages.AntCommandMessage{
		SYNC: buffer[0],
		Id:   buffer[2],
		Name: name,
		Data: data,
	}

	log.Println("reply received correcly")
	return reply, nil
}
