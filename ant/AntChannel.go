package ant

import (
	"github.com/pjvds/antport/messages"
)

// Channel types
const (
	BIDIRECTIONAL_RECEIVE  = 0x00
	BIDIRECTIONAL_TRANSMIT = 0x10

	SHARED_BIDIRECTIONAL_RECEIVE  = 0x20
	SHARED_BIDIRECTIONAL_TRANSMIT = 0x30

	UNIDIRECTIONAL_RECEIVE_ONLY  = 0x40
	UNIDIRECTIONAL_TRANSMIT_ONLY = 0x50
)

/* ANT channel consists of one or more transmitting 
   nodes and one or more receiving nodes depending on 
   the network topology. Any node can transmit or 
   receive so the channels are bi-directional. */
type AntChannel struct {
	ant *AntContext

	// The channel number range acceptable values being
	// from 0 to the maximum number defined by the ANT 
	// implementation.
	number byte
}

func (channel AntChannel) SetNetworkKey(networkNumber byte, key [8]byte) {
	ant := channel.ant
	cmd := messages.CreateSetNetworkKeyCommand(networkNumber, key)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}

func (channel AntChannel) Assign(channelType, networkNumber byte) {
	ant := channel.ant
	cmd := messages.CreateAssignChannelCommand(channel.number, channelType, networkNumber)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}

func (channel AntChannel) SetId(deviceNumber int, networkNumber byte, transType byte) {
	ant := channel.ant
	cmd := messages.CreateSetChannelIdCommand(channel.number, deviceNumber,
		networkNumber, transType)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}

func (channel AntChannel) SetPeriod(period uint16) {
	ant := channel.ant
	cmd := messages.CreateSetChannelPeriodCommand(channel.number, period)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}

func (channel AntChannel) SetSearchTimeout(timeout byte) {
	ant := channel.ant
	cmd := messages.CreateSetChannelSearchTimeoutCommand(channel.number, timeout)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}

func (channel AntChannel) SetRfFrequenty(rfFrequenty byte) {
	ant := channel.ant
	cmd := messages.CreateSetChannelRfFrequentyCommand(channel.number, rfFrequenty)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}

func (channel AntChannel) SetSearchWaveform(waveform uint16) {
	ant := channel.ant
	cmd := messages.CreateSetChannelSearchWaveformCommand(channel.number, waveform)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}

func (channel AntChannel) Open() {
	ant := channel.ant
	cmd := messages.CreateOpenChannelCommand(channel.number)

	ant.SendCommand(cmd)
	ant.ReceiveReply()
}
