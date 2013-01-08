package antport

import (
	"log"
	"testing"
	"time"
)

// func TestAntChannelClose(t *testing.T) {
// 	channel, ctx := GetSingleChannelOrFail(t)
// 	defer ctx.Close()

// 	channel.Close()

// 	if channel.device != nil {
// 		t.Error("close doesn't cleanup device")
// 	}
// }

// func TestSendBecon(t *testing.T) {
// 	channel, ctx := GetSingleChannelOrFail(t)
// 	defer ctx.Close()
// 	defer channel.Close()

// 	log.Println("creating writer...")
// 	writer, err := channel.CreateWriter()

// 	for err != nil {
// 		log.Printf("read was not created: %s", err)

// 		writer, err = channel.CreateWriter()
// 	}

// 	log.Println("writer created!!!")
// 	log.Println("writing...")

// 	buffer := make([]byte, 8)
// 	written, err := writer.Write(buffer)
// 	for written == 0 {
// 		log.Printf("no bytes written: %s", err)

// 		written, err = writer.Write(buffer)
// 	}

// 	log.Println("bytes written!!!")
// }

func TestWaitForBecon(t *testing.T) {
	channel, ctx := GetSingleChannelOrFail(t)
	defer ctx.Close()
	defer channel.Close()

	log.Println("creating reader...")
	reader, err := channel.CreateReader()

	for err != nil {
		log.Printf("read was not created: %s", err)
		time.Sleep(10000)

		reader, err = channel.CreateReader()
	}

	log.Println("reader created!!!")
	log.Println("reading...")

	buffer := make([]byte, 8)
	read, err := reader.Read(buffer)
	for read == 0 {
		log.Printf("no bytes read: %s", err)
		time.Sleep(10000)

		read, err = reader.Read(buffer)
	}

	log.Printf("bytes read!!!: %v\n", string(buffer[0:read]))
}

// func TestCommunication(t *testing.T) {
// 	channel, ctx := GetSingleChannelOrFail(t)
// 	defer ctx.Close()
// 	defer channel.Close()

// 	channel.SendAck()
// 	msg := channel.ReceiveMessage()

// 	if msg == nil {
// 		t.Error("messages not created")
// 	}
// }

func GetSingleChannelOrFail(t *testing.T) (*AntUsbDevice, *AntContext) {
	ctx := NewContext()

	channels, err := ctx.ListAntUsbDevices()

	if err != nil {
		t.Errorf("ListChannels failed with error: %v", err)
	}

	if len(channels) == 0 {
		CloseContextAndFail(t, ctx, "no channels available")
	}

	if len(channels) > 1 {
		CloseContextAndFail(t, ctx, "multiple channels available")
	}

	channel := channels[0]
	return channel, ctx
}

func CloseContextAndFail(t *testing.T, ctx *AntContext, message string) {
	defer ctx.Close()
	t.Error("No channels available")
	t.FailNow()
}
