package main

import "fmt"

func Sqrt(x float64) float64 {
	// TODO: Consider decimals and if they are UP or DOWN
	loopEnd := true
	var finalResult float64
	for z := 1; loopEnd; z++ {
		pow := float64(z * z)
		if pow > x {
			finalResult = float64(z - 1)
			loopEnd = false
		}
	}
	return finalResult
}

func main() {
	fmt.Printf("%f\n", Sqrt(1))
	fmt.Printf("%f\n", Sqrt(2))
	fmt.Printf("%f\n", Sqrt(4))
	fmt.Printf("%f\n", Sqrt(120))
}
