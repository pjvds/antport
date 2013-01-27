package ant

const (
	MESG_REQUEST_MESSAGE_ID = 0x4D
)

func RequestMessage(channelNumber byte, msgId byte) AntMessage {
	return NewAntMessage(MESG_TX_SYNC, MESG_REQUEST_MESSAGE_ID, []byte{channelNumber, msgId})
}
