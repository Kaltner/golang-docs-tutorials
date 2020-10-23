package main

import "testing"

func Test_fibonnaci(t *testing.T) {
	sequence := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

	f := fibonacci()
	for _, value := range sequence {
		fibValue := f()
		if fibValue != value {
			t.Errorf("Unexpected number generated in the sequence. \n Expected: %d \n Got: %d", value, fibValue)
		}
	}
}
