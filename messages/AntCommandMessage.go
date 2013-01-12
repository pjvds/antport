package messages

const (
	DIR_IN  = "IN"
	DIR_OUT = "OUT"
)

type AntCommandMessage struct {
	SYNC      byte
	Direction string
	Id        byte
	Name      string
	Data      []byte
}

// Creates a new AntCommandInfo message
func newMessage(direction string, id byte, name string, data []byte) *AntCommandMessage {
	return &AntCommandMessage{
		Direction: direction,
		Id:        id,
		Name:      name,
		Data:      data,
	}
}
