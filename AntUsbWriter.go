package antport

import ()

type AntUsbWriter interface {
	Write(buffer []byte) (n int, err error)
}
