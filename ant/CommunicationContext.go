package ant

import (
	"code.google.com/p/log4go"
	"container/list"
	"github.com/pjvds/antport/hardware"
	"sync"
)

type CommunicationContext struct {
	// Channel for incomming messages. In other words
	// messages that come from the AntDevice.
	input chan AntMessage

	// Channel for outgoing messages. In other words
	// messages that are written to the AntDevice.
	Output chan SendMessageTicket

	device hardware.AntDevice

	receiver MessageReceiver
	sender   MessageSender

	isOpen     bool
	isOpenLock sync.Mutex

	communicating sync.WaitGroup

	waitingTickets *list.List
	unmatchedInput *list.List

	matchLock sync.Mutex
}

func NewCommunicationContext(device hardware.AntDevice) CommunicationContext {
	receiver := newReceiver(device)
	sender := newSender(device)

	return CommunicationContext{
		input:  make(chan AntMessage, 255),
		Output: make(chan SendMessageTicket, 255),

		device:   device,
		receiver: receiver,
		sender:   sender,

		waitingTickets: list.New(),
		unmatchedInput: list.New(),
	}
}

func (ctx *CommunicationContext) Open() {
	log4go.Debug("opening communiction...")

	ctx.isOpenLock.Lock()
	defer ctx.isOpenLock.Unlock()

	if !ctx.isOpen {
		ctx.device.Reset()
		ctx.isOpen = true

		go ctx.readLoop()
		go ctx.writeLoop()
		go ctx.matchLoop()

		log4go.Debug("communication opened")
	} else {
		log4go.Debug("communication already open, nothing changed")
	}
}

func (ctx *CommunicationContext) readLoop() {
	ctx.communicating.Add(1)
	defer ctx.communicating.Done()

	log4go.Debug("read loop started")

	for ctx.isOpen {
		msg, err := ctx.receiver.Receive()

		if err != nil {
			log4go.Warn("error while receiving from device: %v", err.Error())
		} else {
			ctx.input <- *msg
		}
	}

	log4go.Debug("read loop finished")
}

func (ctx *CommunicationContext) Send(msg AntMessage) *SendMessageTicket {
	ticket := SendMessageTicket{ctx: ctx, msg: msg, send: make(chan AntMessage, 1), error: make(chan error)}
	ctx.Output <- ticket

	return &ticket
}

func (ctx *CommunicationContext) registerWaitForReply(msg AntMessage, matcher func(AntMessage) bool) WaitForReplyTicket {
	ticket := WaitForReplyTicket{
		msg:     msg,
		matcher: matcher,
		reply:   make(chan AntMessage, 1),
		error:   make(chan error, 1),
	}

	ctx.matchLock.Lock()
	defer ctx.matchLock.Unlock()
	var match bool

	for e := ctx.unmatchedInput.Front(); e != nil; e.Next() {
		msg := e.Value.(AntMessage)
		if matcher(msg) {
			ctx.unmatchedInput.Remove(e)
			match = true
			ticket.reply <- msg
			break
		}
	}

	if !match {
		ctx.waitingTickets.PushBack(ticket)
	}
	return ticket
}

func (ctx *CommunicationContext) matchLoop() {
	ctx.communicating.Add(1)
	defer ctx.communicating.Done()

	log4go.Debug("match loop started")

	for ctx.isOpen {
		input := <-ctx.input
		var found bool

		log4go.Debug("new message received in match loop, matching...")
		ctx.matchLock.Lock()

		nWaiting := ctx.waitingTickets.Len()
		log4go.Debug("There are %v waiting tickets.", nWaiting)

		if nWaiting > 0 {
			for e := ctx.waitingTickets.Front(); e != nil; e = e.Next() {
				waitTicket := e.Value.(WaitForReplyTicket)
				isMatch := waitTicket.matcher(input)

				if isMatch {
					log4go.Debug("Match found!")
					waitTicket.reply <- input
					ctx.waitingTickets.Remove(e)
					found = true
					break
				}
			}
		}

		// If there was no match, push this message
		// to the unmatched list. This list is checked
		// when a new match is registered.
		if !found {
			log4go.Debug("No match found. Pushing message to unmatched input list.")
			ctx.unmatchedInput.PushBack(input)
		}

		ctx.matchLock.Unlock()
	}
}

func (ctx *CommunicationContext) writeLoop() {
	ctx.communicating.Add(1)
	defer ctx.communicating.Done()

	log4go.Debug("write loop started")

	for ctx.isOpen {
		tckt, ok := <-ctx.Output

		if ok {
			log4go.Debug("found new output in output channel")
			err := ctx.sender.Send(tckt.msg)

			if err != nil {
				log4go.Warn("error while sending to device: %v", err.Error())
				tckt.error <- err
			} else {
				tckt.isSend = true
				tckt.send <- tckt.msg
			}
		} else {
			log4go.Warn("output channel closed")
		}
	}

	log4go.Debug("write loop finished")
}

func (ctx *CommunicationContext) Close() {
	log4go.Debug("closing communication...")

	ctx.isOpenLock.Lock()
	defer ctx.isOpenLock.Unlock()

	if ctx.isOpen {
		ctx.isOpen = false
		close(ctx.input)
		close(ctx.Output)

		ctx.communicating.Wait()
		ctx.device.Close()

		ctx.device = nil
		ctx.input = nil
		ctx.Output = nil

		log4go.Debug("communication closed")
	} else {
		log4go.Debug("communication already closed, nothing changed")
	}
}
