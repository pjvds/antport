package ant

import (
	"code.google.com/p/log4go"
)

type AntContext struct {
	communication *CommunicationContext
}

func NewAntContext(communication *CommunicationContext) AntContext {
	ctx := AntContext{
		communication: communication,
	}

	return ctx
}

func (ctx *AntContext) Init() error {
	ctx.communication.Open()
	ticket := ctx.communication.Send(RequestMessage(0, MESG_CAPABILITIES_ID))
	response, err := ticket.WaitForReply(IsCapabilities)

	if err != nil {
		log4go.Warn("AntContext didn't initialize succesfully: %v", err.Error())
		return err
	}

	response.AsCapabilities()
	log4go.Debug("AntContext initilized succesfully")
	return nil
}

func (ctx *AntContext) initializeCapabilities() {

}

func (ctx *AntContext) Close() {
	ctx.communication.Close()
	ctx.communication = nil
}
