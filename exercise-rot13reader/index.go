package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := range b {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] += 13
			if b[i] > 'z' {
				b[i] = b[i] - 'z' + 'a' - 1
			}
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 13
			if b[i] > 'Z' {
				b[i] = b[i] - 'Z' + 'A' - 1
			}
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
