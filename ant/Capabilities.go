package ant

type Capabilities struct {
	MaxChannels        byte
	MaxNetworks        byte
	StandardOperations byte
	AdvancedOperations byte
}

func (msg *AntMessage) AsCapabilities() Capabilities {
	if msg.Id != MESG_CAPABILITIES_ID {
		panic("Cannot create Capabilities structure from AntMessage with other id then " + string(MESG_CAPABILITIES_ID))
	}

	if msg.Length < 4 {
		panic("Invalid message lenght: " + string(msg.Length))
	}

	return Capabilities{
		MaxChannels:        msg.Data[0],
		MaxNetworks:        msg.Data[1],
		StandardOperations: msg.Data[2],
		AdvancedOperations: msg.Data[3],
	}
}

func IsCapabilities(msg AntMessage) bool {
	return msg.Id == MESG_CAPABILITIES_ID
}
