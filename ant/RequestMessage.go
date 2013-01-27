package ant

func RequestMessage(channelNumber byte, msgId byte) AntMessage {
	return NewAntMessage(MESG_TX_SYNC, MESG_REQUEST_MESSAGE_ID, []byte{channelNumber, msgId})
}
