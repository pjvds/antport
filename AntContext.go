package antport

import (
	"log"
)

const (
	DIR_IN  = "IN"
	DIR_OUT = "OUT"
)

type AntContext struct {
	device       AntDevice
	Initialized  bool
	Capabilities *AntCapabilityInfo
	MaxRetry     int
}

type AntCapabilityInfo struct {
}

func CreateAntContext(device AntDevice) *AntContext {
	return &AntContext{
		device:   device,
		MaxRetry: 5,
	}
}

func (ctx *AntContext) Init() {
	ctx.ResetSystem()
}

func (ctx *AntContext) ResetSystem() {
	ctx.SendCommand(CreateResetCommand())
	ctx.ReceiveReply()

	ctx.SendCommand(CreateRequestMessageCommand(0, 0x54))
	ctx.ReceiveReply()
}

func (ctx *AntContext) SendCommand(cmd *AntCommand) {
	data := cmd.Pack()
	ctx.device.Write(data)
}

func (ctx *AntContext) ReceiveReply() (reply *AntCommand, err error) {
	buffer := make([]byte, 8)
	n, err := ctx.device.Read(buffer)

	for retries := 1; retries < ctx.MaxRetry; retries++ {
		if err != nil {
			log.Println("error while receiving reply: " + err.Error())
			log.Printf("will retry (%v/%v)", retries, ctx.MaxRetry)

			n, err = ctx.device.Read(buffer)
		}
	}

	if err != nil {
		return nil, err
	}

	return newMessage(DIR_IN, 0x00, "RAWCOMMAND", buffer[0:n]), nil
}

// Creates a new AntCommand message
func newMessage(direction string, id byte, name string, data []byte) *AntCommand {
	return &AntCommand{
		Direction: direction,
		Id:        id,
		Name:      name,
		Data:      data,
	}
}
