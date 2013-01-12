package antport

type AntFsContext struct {
	ant     *AntContext
	channel *AntChannel
	network *AntNetwork
}

func NewAntFsContext(ant *AntContext) *AntFsContext {
	return &AntFsContext{
		ant: ant,
	}
}

func (ctx *AntFsContext) OpenAntsFsSearchChannel() {
	ctx.ant.ResetSystem()
	ctx.channel = ctx.ant.Channels[0]
	ctx.network = ctx.ant.Networks[0]

	channel := ctx.channel
	channel.SetNetworkKey(0, SEARCH_NETWORK_KEY)
	channel.Assign(0x00, ctx.network.number)
	channel.SetId(0, 0, 0)
	channel.Open()
}
