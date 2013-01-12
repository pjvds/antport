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
}
