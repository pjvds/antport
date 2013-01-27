package hardware

type AntDevice interface {
	Reset()
	Read(buffer []byte) (n int, err error)
	Write(data []byte) (n int, err error)
	Close()
}
