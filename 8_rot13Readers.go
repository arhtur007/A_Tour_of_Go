package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, e error) {
	n, e = r.r.Read(b)
	for i, c := range b {

		switch {
		case (c >= 'a' && c <= 'm') || (c >= 'A' && c <= 'M'):
			b[i] += 13
		case (c > 'm' && c <= 'z') || (c > 'M' && c <= 'Z'):
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
