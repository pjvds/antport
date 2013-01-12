package messages

import ()

const (
	SetChannelPeriodCommandId   = byte(0x43)
	SetChannelPeriodCommandName = "SET_CHANNEL_PERIOD"
)

type SetChannelPeriodCommand struct {
	AntCommandInfo

	ChannelNumber byte

	/* The channel period represents the basic message rate of
	data packets sent by the master. By default a broadcast 
	data packet will be sent or received on every timeslot at 
	this rate.  The channel message rate can range from 0.5Hz
	to above 200Hz with the upper limit being implementation 
	specific.  The default message rate is 4Hz, which is chosen 
	to provide good performance as described below.  It is 
	recommended that the message rate be left at the default
	to provide more readily discoverable networks
	with good performance characteristics. */
	MessagingPeriod byte
}

func CreateSetChannelPeriodCommand(channelNumber byte, messagingPeriod byte) SetChannelPeriodCommand {
	cmd := newAntCommandInfo(SetChannelIdCommandId, SetChannelPeriodCommandName)
	return SetChannelPeriodCommand{
		AntCommandInfo:  cmd,
		ChannelNumber:   channelNumber,
		MessagingPeriod: messagingPeriod,
	}
}

func (cmd SetChannelPeriodCommand) Data() []byte {
	return []byte{
		cmd.ChannelNumber,
		cmd.MessagingPeriod,
	}
}
