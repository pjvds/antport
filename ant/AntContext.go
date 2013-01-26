package ant

import (
	"github.com/pjvds/antport/hardware"
	"github.com/pjvds/antport/messages"
	"log"
)

type AntContext struct {
	device       hardware.AntDevice
	sender       messageSender
	receiver     messageReceiver
	communicator AntCommunicator

	Initialized  bool
	Capabilities *AntCapabilityInfo
	MaxRetry     int
	Channels     []*AntChannel
	Networks     []*AntNetwork
}

func CreateAntContext(device hardware.AntDevice) *AntContext {
	sender := newSender(device)
	receiver := newReceiver(device)

	communicator := newCommunicator(receiver, sender)

	return &AntContext{
		device:       device,
		sender:       sender,
		receiver:     receiver,
		communicator: communicator,
		MaxRetry:     1000,
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

func (ctx *AntContext) SendCommand(message messages.AntCommand) {
	ctx.sender.SendCommand(message)
}

func (ctx *AntContext) ReceiveReply() (reply *messages.AntCommandMessage, err error) {
	return ctx.ReceiveReply()
}

// A hard reset can be preformed on ANT hardware by
// sending 15 zero's. This method will retry until succeeds.
func (ctx *AntContext) HardResetSystem(retryCount byte) error {
	log.Println("hard resetting device system")
	data := make([]byte, 15)
	n, err := ctx.device.Write(data)

	for retry := byte(0); n != 15 || err != nil; retry++ {
		if retry > retryCount {
			return err
		}

		log.Println("hard reset failed.")
		n, err = ctx.device.Write(data)
	}

	log.Println("hard reset ok")
	return nil
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
