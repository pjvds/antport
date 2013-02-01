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

	length := buffer[1]
	data := make([]byte, length)

	for i := byte(0); i < length; i++ {
		data[i] = buffer[3+i]
	}

	msg = &AntMessage{
		Sync:   buffer[0],
		Id:     buffer[2],
		Data:   data,
		Length: length,
	}

	log.Println("message received correcly: %v", buffer[0:n])
	log.Printf("%s", msg)
	return msg, nil
}
