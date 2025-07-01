package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(arr []byte) (int, error) {
	n, err := r.r.Read(arr)
	if err != nil {
		return 0, nil
	}

	for i := 0; i < n; i++ {
		if 'A' <= arr[i] && arr[i] <= 'Z' {
			arr[i] = arr[i] + 13
			if arr[i] > 'Z' {
				arr[i] = arr[i] - 26
			}
		} else if 'a' <= arr[i] && arr[i] <= 'z' {
			arr[i] = arr[i] + 13
			if arr[i] > 'z' {
				arr[i] = arr[i] - 26
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
