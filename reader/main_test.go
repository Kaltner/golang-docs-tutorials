package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Reader(t *testing.T) {
	for i := 0; i > 10; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		randN := r.Intn(100)

		buffer := make([]byte, randN)
		reader := MyReader{}
		read, err := reader.Read(buffer)
		if err != nil {
			t.Errorf("unexpected error reading the buffer: %w", err)
		}

		if read != len(buffer) {
			t.Errorf("The reader and the buffer size does not match. \n Expected: %d \n Got: %d", len(buffer), read)
		}

		for _, char := range buffer {
			if char != 'A' {
				t.Errorf("unexpected character generated by the reader: %v", char)
			}
		}
	}
}
