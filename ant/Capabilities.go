package ant

func RequestMessage(channelNumber byte, msgId byte) AntMessage {
	return NewAntMessage(0x00, 0x4D, []byte{channelNumber, msgId})
}
