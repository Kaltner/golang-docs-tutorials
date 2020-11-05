package main

import (
	"fmt"
	"os"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

// This is just for testing and understanding reading
func fileReader() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	// _ = fileinfo.Size()
	buffer := make([]byte, filesize)

	fmt.Println("file bytes before read: ", buffer)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bytes read: ", bytesread)
	fmt.Println("file bytes: ", buffer)
	fmt.Println("bytestream to string: ", string(buffer))
}

func main() {
	fileReader()
	// reader := MyReader{}
	// read, err := reader.Read([]byte{0})
	// fmt.Printf("%q, %w", read, err)
}
