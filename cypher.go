package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
	distance byte
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	count, err := rot.r.Read(p)

	for i :=0; i < count; i++ {
		if p[i] >= 97 {
			p[i] = ((p[i] - 97) + rot.distance) % 26 + 97
		}
	}

	return count, err
}

func main() {
	s := strings.NewReader("lbh penpxrq gur pbqr!")
	r := rot13Reader{s, 13}
	io.Copy(os.Stdout, &r)
}
