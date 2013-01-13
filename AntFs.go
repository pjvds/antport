package antport

type AntFsContext struct {
	ant     *AntContext
	channel *AntChannel
	network *AntNetwork
}

const (
	search_freq     = 0x32
	search_period   = 0x1000
	search_timeout  = 255
	search_waveform = 0x5300
)

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
	channel.SetPeriod(4096)
	channel.SetRfFrequenty(50)
	channel.SetSearchTimeout(255)
	channel.SetSearchWaveform(search_waveform)
	channel.SetId(0, 0x01, 0)
	channel.Open()
}
