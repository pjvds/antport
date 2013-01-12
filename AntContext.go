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
	Channels     []*AntChannel
	Networks     []*AntNetwork
}

type AntNetwork struct {
	ant *AntContext

	// The Network Number is an 8-bit field with the 
	// range of acceptable values being from 0 to the 
	// maximum number defined by the ANT implementation.
	number byte

	// The Network Key is an 8-byte field which is configurable
	// by the host application. A particular Network
	// Number will have a corresponding Network Key.  
	// The Network Number and the Network Key together provide 
	// the ability to deploy a network with varied levels of 
	// access control and security options. 
	key [8]byte
}

type AntCapabilityInfo struct {
	MaxChannels byte
	MaxNetworks byte
}

func CreateAntContext(device AntDevice) *AntContext {
	return &AntContext{
		device:   device,
		MaxRetry: 5,
	}
}

// Initialize the context
func (ctx *AntContext) Init() {
	ctx.ResetSystem()
}

// Reset system and initialize capabilities
func (ctx *AntContext) ResetSystem() {
	ctx.SendCommand(CreateResetCommand())
	ctx.ReceiveReply()

	ctx.initCapabilities()
}

// A hard reset can be preformed on ANT hardware by
// sending 15 zero's. This method will retry until succeeds.
func (ctx *AntContext) HardResetSystem() {
	log.Println("hard resetting device system")
	data := make([]byte, 15)
	n, err := ctx.device.Write(data)

	for n != 15 || err != nil {
		log.Println("hard reset failed.")
		n, err = ctx.device.Write(data)
	}

	log.Println("hard reset ok")
}

func (ctx *AntContext) initCapabilities() {
	ctx.SendCommand(CreateRequestMessageCommand(0, 0x54))
	reply, err := ctx.ReceiveReply()

	if err != nil {
		log.Println("error while requesting capabilities: " + err.Error())
	}

	ctx.Capabilities = &AntCapabilityInfo{
		MaxChannels: reply.Data[0],
		MaxNetworks: reply.Data[1],
	}

	// Create channels
	channels := make([]*AntChannel, ctx.Capabilities.MaxChannels)
	for i := 0; i < len(channels); i++ {
		channels[i] = &AntChannel{
			ant:    ctx,
			number: byte(i),
		}
	}
	ctx.Channels = channels

	// Create networks
	networks := make([]*AntNetwork, ctx.Capabilities.MaxNetworks)
	for i := 0; i < len(networks); i++ {
		var key [8]byte
		networks[i] = &AntNetwork{
			ant:    ctx,
			number: byte(i),
			key:    key,
		}
	}
	ctx.Networks = networks

	log.Printf("context capabilities initialized: %s", ctx.Capabilities)
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
