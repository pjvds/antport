package ant

import (
	"testing"
)

func TestRequestMessage(t *testing.T) {
	channelNumber := byte(0x01)
	msgId := byte(0x42)

	msg := RequestMessage(channelNumber, msgId)

	if msg.Id != MESG_REQUEST_MESSAGE_ID {
		t.Errorf("Id not set correctly. Expected %v actual %v", MESG_REQUEST_MESSAGE_ID, msg.Id)
	}
	if msg.Length != 2 {
		t.Errorf("Length not set correctly. Expected %v actual %v", 2, msg.Length)
	}
	if msg.Data[0] != channelNumber {
		t.Errorf("Channel number data not set correctly. Expected %v actual %v", channelNumber, msg.Data[0])
	}

	if msg.Data[1] != msgId {
		t.Errorf("Message id data not set correctly. Expected %v actual %v", msgId, msg.Data[1])
	}
}
