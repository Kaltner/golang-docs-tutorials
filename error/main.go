package main

import (
	"fmt"
	"strconv"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("parameter cannot be negative: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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
		-1,
		-100,
	}
	for _, n := range numbers {
		result, err := Sqrt(n)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%f\n", result)
	}
}
