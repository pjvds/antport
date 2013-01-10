package antport

import ()

type AntUsbReader interface {
	Read(buffer []byte) (n int, err error)
}
