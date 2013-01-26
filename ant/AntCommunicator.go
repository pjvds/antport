package ant

import (
	"github.com/pjvds/antport/messages"
)

type AntCommunicator struct {
	receiver messageReceiver
	sender   messageSender

	outbox chan MessageTicket
}

func newCommunicator(receiver messageReceiver, sender messageSender) AntCommunicator {
	return AntCommunicator{
		receiver: receiver,
		sender:   sender,
		outbox:   make(chan MessageTicket, 255),
	}
}

func (c *AntCommunicator) Send(message messages.AntCommand) MessageTicket {

	ticket := newMessageTicket(message)
	c.outbox <- ticket

	ticket = <-c.outbox
	ok, err := c.sender.SendCommand(message)

	if ok {
		ticket.onSend <- message
	} else {
		ticket.onError <- err
	}

	return ticket

}
