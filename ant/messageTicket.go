package ant

import (
	"github.com/pjvds/antport/messages"
)

type MessageTicket struct {
	communicator *AntCommunicator
	payload      messages.AntCommand
	send         bool

	onError chan error
	onSend  chan messages.AntCommand
	onReply chan messages.AntCommand
}

func newMessageTicket(communicator *AntCommunicator, message messages.AntCommand) MessageTicket {
	return MessageTicket{
		communicator: communicator,
		payload:      message,
		onError:      make(chan error, 1),
		onSend:       make(chan messages.AntCommand, 1),
		onReply:      make(chan messages.AntCommand, 1),
	}
}

func (ticket MessageTicket) WaitForCompletion() (err error) {
	select {
	case err = <-ticket.onError:
		return err
	case <-ticket.onSend:
		return nil
	}

	panic("missing case statement in WaitForCompletion")
}

func (ticket MessageTicket) WaitForReply() (reply messages.AntCommand, err error) {

	return nil, nil

	panic("missing case statement in WaitForReply")
}
