package ant

import (
	"github.com/pjvds/antport/messages"
	"sync"
)

type AntCommunicator struct {
	receiver messageReceiver
	sender   messageSender

	outbox         chan MessageTicket
	communications sync.WaitGroup
	stopping       bool
}

func newCommunicator(receiver messageReceiver, sender messageSender) AntCommunicator {
	return AntCommunicator{
		receiver: receiver,
		sender:   sender,
		outbox:   make(chan MessageTicket, 255),
	}
}

func (c *AntCommunicator) Start() {
	c.communications.Wait()
	c.communications.Add(1)
	defer c.communications.Done()

	go c.process()
}

func (c *AntCommunicator) Stop() {
	c.stopping = true
	close(c.outbox)

	c.communications.Wait()
}

func (c *AntCommunicator) process() {
	c.communications.Add(1)
	defer c.communications.Done()

	for !c.stopping {
		ticket := <-c.outbox

		ok, err := c.sender.SendCommand(ticket.payload)

		if ok {
			ticket.onSend <- ticket.payload
		} else {
			ticket.onError <- err
		}
	}
}

func (c *AntCommunicator) Send(message messages.AntCommand) MessageTicket {
	ticket := newMessageTicket(message)
	c.outbox <- ticket

	return ticket
}
