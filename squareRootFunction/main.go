package main

import (
	"fmt"
	"strconv"
)

func Sqrt(x float64) (float64, error) {
	var lastIteration string
	repetitionTicks := 0
	z := 1.0
	for repetitionTicks < 3 {
		z -= (z*z - x) / (2 * z)
		zDecimal := fmt.Sprintf("%.2f", z)
		if zDecimal != lastIteration {
			lastIteration = zDecimal
		} else {
			repetitionTicks++
		}
	}
	finalResult, err := strconv.ParseFloat(lastIteration, 64)
	if err != nil {
		return 0, err
	}
	return finalResult, nil
}

func main() {
	numbers := []float64{
		1,
		1.5,
		2,
		3,
		4,
		5,
		120,
	}
	for _, n := range numbers {
		result, err := Sqrt(n)
		if err != nil {
			fmt.Sprintf("%v", err)
			break
		}
		fmt.Printf("%f\n", result)
	}
}
