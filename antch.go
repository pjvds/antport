package antport

import (
	"github.com/kylelemons/gousb/usb"
)

const (
	ANT_VENDOR_ID  = 0xfcf
	ANT_PRODUCT_ID = 0x1008
)

type AntChannel struct {
	device *usb.Device
}
