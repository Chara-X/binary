package binary

import "io"

type Writer struct {
	io.Writer
	bits  uint64
	nbits int
}

func (w *Writer) WriteBits(b uint64, n int) {
	w.bits |= b << w.nbits
	w.nbits += n
	for w.nbits >= 8 {
		w.Write([]byte{byte(w.bits)})
		w.bits >>= 8
		w.nbits -= 8
	}
}

func (w *Writer) Close() {
	if w.nbits > 0 {
		w.Write([]byte{byte(w.bits)})
	}
}
