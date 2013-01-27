package ant

import (
	"code.google.com/p/log4go"
	"github.com/pjvds/antport/hardware"
	"sync"
)

type CommunicationContext struct {
	// Channel for incomming messages. In other words
	// messages that come from the AntDevice.
	Input chan AntMessage

	// Channel for outgoing messages. In other words
	// messages that are written to the AntDevice.
	Output chan AntMessage

	device hardware.AntDevice

	receiver MessageReceiver
	sender   MessageSender

	clossing      bool
	communicating sync.WaitGroup
}

func NewCommunicationContext(device hardware.AntDevice) CommunicationContext {
	receiver := newReceiver(device)
	sender := newSender(device)

	return CommunicationContext{
		Input:  make(chan AntMessage, 255),
		Output: make(chan AntMessage, 255),

		device:   device,
		receiver: receiver,
		sender:   sender,
	}
}

func (ctx *CommunicationContext) Open() {
	ctx.communicating.Add(1)
	defer ctx.communicating.Done()

	ctx.device.Reset()
	go ctx.readLoop()
	go ctx.writeLoop()
}

func (ctx *CommunicationContext) readLoop() {
	ctx.communicating.Add(1)
	defer ctx.communicating.Done()

	log4go.Debug("read loop started")

	for !ctx.clossing {
		msg, err := ctx.receiver.Receive()

		if err != nil {
			log4go.Warn("error while receiving from device: %v", err.Error())
		} else {
			ctx.Input <- *msg
		}
	}

	log4go.Debug("read loop finished")
}

func (ctx *CommunicationContext) writeLoop() {
	ctx.communicating.Add(1)
	defer ctx.communicating.Done()

	log4go.Debug("write loop started")

	for !ctx.clossing {
		msg, ok := <-ctx.Output

		if ok {
			log4go.Debug("found new output in output channel")
			err := ctx.sender.Send(msg)

			if err != nil {
				log4go.Warn("error while sending to device: %v", err.Error())
			}
		} else {
			log4go.Warn("output channel closed")
		}
	}

	log4go.Debug("write loop finished")
}

func (ctx *CommunicationContext) Close() {
	ctx.clossing = true
	close(ctx.Input)
	close(ctx.Output)

	ctx.communicating.Wait()
	ctx.device.Close()

	ctx.device = nil
	ctx.Input = nil
	ctx.Output = nil
}
