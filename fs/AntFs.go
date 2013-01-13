package fs

import (
	"github.com/pjvds/antport/ant"
)

type AntFsContext struct {
	ant     *ant.AntContext
	channel *ant.AntChannel
	network *ant.AntNetwork
}

var (
	search_network_key = [8]byte{
		0xa8, 0xa4, 0x23, 0xb9,
		0xf5, 0x5e, 0x63, 0xc1}
)

const (
	search_freq     = 0x32
	search_period   = 0x1000
	search_timeout  = 255
	search_waveform = 0x5300
)

func NewAntFsContext(antCtx *ant.AntContext) *AntFsContext {
	return &AntFsContext{
		ant: antCtx,
	}
}

func (ctx *AntFsContext) OpenAntsFsSearchChannel() {
	ctx.ant.ResetSystem()
	ctx.channel = ctx.ant.Channels[0]
	ctx.network = ctx.ant.Networks[0]

	network := ctx.network
	network.SetNetworkKey(search_network_key)

	channel := ctx.channel
	channel.Assign(0x00, ctx.network.Number)
	channel.SetPeriod(4096)
	channel.SetRfFrequenty(50)
	channel.SetSearchTimeout(255)
	channel.SetSearchWaveform(search_waveform)
	channel.SetId(0, 0x01, 0)
	channel.Open()
}
