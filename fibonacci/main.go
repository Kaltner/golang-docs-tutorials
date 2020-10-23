package main

import "fmt"

func fibonacci() func() int {
	sequence := make([]int, 0)
	return func() int {
		l := len(sequence)

		var nextN int
		if l == 0 {
			nextN = 0
		}
		if l == 1 {
			nextN = 1
		}
		if l >= 2 {
			nextN = sequence[l-1] + sequence[l-2]
		}
		sequence = append(sequence, nextN)
		return sequence[l] // Will return the latest value, since sequence is 0 indexed
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
