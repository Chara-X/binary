package binary

import "io"

type Reader struct {
	io.Reader
}

func (r *Reader) ReadBits(n uint8) (u uint64, err error) {
	panic("implement me")
}
