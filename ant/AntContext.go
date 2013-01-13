package ant

import (
	"fmt"
	"github.com/pjvds/antport/hardware"
	"github.com/pjvds/antport/messages"
	"log"
)

type AntContext struct {
	device       hardware.AntDevice
	Initialized  bool
	Capabilities *AntCapabilityInfo
	MaxRetry     int
	Channels     []*AntChannel
	Networks     []*AntNetwork
}

func CreateAntContext(device hardware.AntDevice) *AntContext {
	return &AntContext{
		device:   device,
		MaxRetry: 1000,
	}
}

// Initialize the context
func (ctx *AntContext) Init() {
	ctx.ResetSystem()
}

// Reset system and initialize capabilities
func (ctx *AntContext) ResetSystem() {
	cmd := messages.CreateResetSystemCommand()

	ctx.SendCommand(cmd)
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
	cmd := messages.CreateRequestMessageCommand(0, 0x54)
	ctx.SendCommand(cmd)
	replyMsg, err := ctx.ReceiveReply()

	if err != nil {
		log.Println("error while requesting capabilities: " + err.Error())
	}

	reply, err := messages.NewCapabilitiesReply(replyMsg)

	if err != nil {
		log.Println("error while creating reply: " + err.Error())
	}

	ctx.Capabilities = &AntCapabilityInfo{
		MaxChannels: reply.MaxChannels,
		MaxNetworks: reply.MaxNetworks,
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
			Number: byte(i),
			Key:    key,
		}
	}
	ctx.Networks = networks

	log.Printf("context capabilities initialized: %s", ctx.Capabilities)
}

func (ctx *AntContext) SendCommand(cmd messages.AntCommand) (ok bool, err error) {
	log.Printf("sending command: %v", cmd.Name())

	msg := messages.NewMessage(cmd)
	data := msg.Pack()
	n, err := ctx.device.Write(data)

	for retries := 1; retries < ctx.MaxRetry+1; retries++ {
		if err != nil {
			log.Println("error while writing to device: " + err.Error())
			log.Printf("will retry (%v/%v)", retries, ctx.MaxRetry)

			n, err = ctx.device.Write(data)
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

	log.Printf("ANT message name: %v", msg.Name)
	return true, nil
}

func (ctx *AntContext) ReceiveReply() (reply *messages.AntCommandMessage, err error) {
	log.Println("receiving reply...")

	buffer := make([]byte, 8)
	n, err := ctx.device.Read(buffer)

	for retries := 1; retries < ctx.MaxRetry+1; retries++ {
		if err != nil {
			log.Println("error while receiving reply: " + err.Error())
			log.Printf("will retry (%v/%v)", retries, ctx.MaxRetry)

			n, err = ctx.device.Read(buffer)
		}
	}

	if err != nil {
		log.Println("error reading from device: " + err.Error())
		return nil, err
	}

	data := make([]byte, 0)
	name := messages.CommandIdToName(buffer[2])
	size := buffer[1]

	log.Printf("ANT message name: %v", name)
	log.Printf("ANT message length: %v", size)
	log.Printf("ANT message raw: %x", buffer[0:n])
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
