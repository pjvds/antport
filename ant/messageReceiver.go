package ant

import (
	"github.com/pjvds/antport/hardware"
	"log"
)

// Receive AntCommand from AntDevice
type MessageReceiver struct {
	hardware.AntDevice

	maxRetry int
}

func newReceiver(device hardware.AntDevice) MessageReceiver {
	return MessageReceiver{
		AntDevice: device,
		maxRetry:  25,
	}
}

func (receiver MessageReceiver) Receive() (msg *AntMessage, err error) {
	buffer := make([]byte, 16)
	n, err := receiver.Read(buffer)

	// for retries := 1; retries < receiver.maxRetry+1; retries++ {
	// 	if err != nil {
	// 		log.Printf("error while receiving message, %v bytes read: %s", n, err)
	// 		log.Printf("will retry (%v/%v)", retries, receiver.maxRetry)

	// 		n, err = receiver.Read(buffer)
	// 	}
	// }

	if err != nil {
		log.Println("error reading from device: " + err.Error())
		return nil, err
	}

	data := make([]byte, 0)
	length := buffer[1]

	if length > 0 {
		data = buffer[3:length]
	}

	msg = &AntMessage{
		Sync:   buffer[0],
		Id:     buffer[2],
		Data:   data,
		Length: length,
	}

	log.Println("message received correcly: %v", buffer[0:n])
	return msg, nil
}
