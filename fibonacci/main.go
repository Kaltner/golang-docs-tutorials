package main

import "fmt"

func fibonacci() func() int {
	sequence := []int{0, 1}
	i := 0
	return func() int {
		if i < 2 {
			i++
			return sequence[i-1]
		}
		l := len(sequence)

		sequence = append(sequence, sequence[l-1]+sequence[l-2])
		return sequence[l] // Will return the latest value, since sequence is 0 indexed
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
