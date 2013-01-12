package antport

type AntFsContext struct {
	ant     *AntContext
	channel *AntChannel
}

func NewAntFsContext(ant *AntContext) *AntFsContext {
	return &AntFsContext{
		ant: ant,
	}
}

func (ctx *AntFsContext) OpenAntsFsSearchChannel() {
	ctx.ant.ResetSystem()
	ctx.channel = ctx.ant.Channels[0]
	ctx.channel.SetId(0, 0, 0)
}
