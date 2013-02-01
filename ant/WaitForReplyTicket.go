package ant

import ()

type WaitForReplyTicket struct {
	msg     AntMessage
	matcher func(AntMessage) (didMatch bool)

	reply chan AntMessage
	error chan error
}
