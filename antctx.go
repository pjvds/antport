package antport

import (
	"github.com/kylelemons/gousb/usb"
)

type AntPortContext struct {
	usb usb.Context
}
