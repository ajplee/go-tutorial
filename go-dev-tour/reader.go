package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(arr []byte) (int, error) {
	size := len(arr)
	for i := 0; i < size; i++ {
		arr[i] = 'A'
	}
	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}
