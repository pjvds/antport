package ant

import (
	"github.com/pjvds/antport/messages"
)

type MessageTicket struct {
	payload messages.AntCommand

	onError chan error
	onSend  chan messages.AntCommand
}

func newMessageTicket(message messages.AntCommand) MessageTicket {
	return MessageTicket{
		payload: message,
		onError: make(chan error, 1),
		onSend:  make(chan messages.AntCommand, 1),
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
