package antport

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

// func (channel *AntChannel) SetId(deviceNumber, networkNumber, transType byte) {
// 	ant := channel.ant
// 	id := &AntChannelId{
// 		deviceNumber:  deviceNumber,
// 		networkNumber: networkNumber,
// 		transType:     transType,
// 	}

// 	ant.SendCommand(CreateSetChannelIdCommand(channel.number, deviceNumber,
// 		deviceTypeId, transType))
// }

type AntChannelId struct {
	deviceNumber  byte
	networkNumber byte
	transType     byte
}
