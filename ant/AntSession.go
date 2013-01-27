package ant

import ()

type SendMessageTicket struct {
	ctx *CommunicationContext
	msg AntMessage

	isSend bool
	send   chan AntMessage
	error  chan error
}

type WaitForReplyTicket struct {
	msg     AntMessage
	matcher func(AntMessage) bool

	reply chan AntMessage
	error chan error
}

func (t *SendMessageTicket) WaitForSendComplete() error {
	select {
	case err := <-t.error:
		return err
	case <-t.send:
		return nil
	}

	panic("missing case statement in WaitForSendComplete!")
}

func (t *SendMessageTicket) WaitForReply(matcher func(AntMessage) bool) {
	t.ctx.registerWaitForReply(t.msg, matcher)
}