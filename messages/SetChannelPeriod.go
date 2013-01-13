package messages

import (
	"bytes"
	"encoding/binary"
)

const (
	SET_CHANNEL_PERIOD_MSG_ID   = byte(0x43)
	SetChannelPeriodCommandName = "SET_CHANNEL_PERIOD"
)

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
type SetChannelPeriodCommand struct {
	AntCommandInfo

	ChannelNumber byte

	/* The channel messaging period in seconds * 32768. 
	Maximum messaging period is ~2 seconds. */
	MessagingPeriod uint16
}

func CreateSetChannelPeriodCommand(channelNumber byte, messagingPeriod uint16) SetChannelPeriodCommand {
	cmd := newAntCommandInfo(SET_CHANNEL_ID_MSG_ID, SetChannelPeriodCommandName)
	return SetChannelPeriodCommand{
		AntCommandInfo:  cmd,
		ChannelNumber:   channelNumber,
		MessagingPeriod: messagingPeriod,
	}
}

func (cmd SetChannelPeriodCommand) Data() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, cmd.ChannelNumber)
	binary.Write(buffer, binary.LittleEndian, cmd.MessagingPeriod)

	return buffer.Bytes()
}
