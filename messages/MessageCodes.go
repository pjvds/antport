package messages

// Messages codes as defined in section 9.5.5
const (
	// Returned on a successful operation 
	RESPONSE_NO_ERROR = byte(0)
	// A receive channel has timed out on searching. The search 
	// is terminated, and the channel has been automatically 
	// closed. In order to restart the search the Open Channel 
	// message must be sent again. 
	EVENT_RX_SEARCH_TIMEOUT = byte(1)

	// A receive channel missed a message which it was 
	// expecting. This would happen when a receiver is tracking 
	// a transmitter and is expecting a message at the set
	// message rate. 
	EVENT_RX_FAIL = byte(2)

	// 	A Broadcast message has been transmitted successfully. 
	// This event should be used to send the next message for 
	// transmission to the ANT device if the node is setup as a 
	// transmitter. 
	EVENT_TX = byte(3)

	// 	A receive transfer has failed. This occurs when a Burst 
	// Transfer Message was incorrectly received. 
	EVENT_TRANSFER_RX_FAILED = byte(4)

	// 	An Acknowledged Data message or a Burst Transfer 
	// sequence has been completed successfully. When 
	// transmitting Acknowledged Data or Burst Transfer, there 
	// is no EVENT_TX message. 
	EVENT_TRANSFER_TX_COMPLETED = byte(5)

	// An Acknowledged Data message, or a Burst Transfer 
	// Message has been initiated and the transmission has
	// failed to complete successfully
	EVENT_TRANSFER_TX_FAILED = byte(6)

	// 	The channel has been successfully closed. When the Host 
	// sends a message to close a channel, it first receives a 
	// RESPONSE_NO_ERROR to indicate that the message was 
	// successfully received by ANT. This event is the actual 
	// indication of the closure of the channel. So, the Host must 
	// use this event message instead of the 
	// RESPONSE_NO_ERROR message to let a channel state 
	// machine continue. 
	EVENT_CHANNEL_CLOSED = byte(7)
)
