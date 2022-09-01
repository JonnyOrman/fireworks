package fireworks

import "io"

type Reader interface {
	Read(ioReader io.Reader) []byte
}
