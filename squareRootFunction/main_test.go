package main

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := map[float64]float64{
		1:   1,
		1.5: 1,
		2:   1,
		3:   1,
		4:   2,
		5:   2,
	}

	for parameter, expectedResult := range cases {
		parameter, expectedResult := parameter, expectedResult
		t.Run(fmt.Sprintf("Test N%f", parameter), func(t *testing.T) {
			// t.Parallel()

			result := Sqrt(parameter)
			if result != expectedResult {
				t.Errorf("Result is not equal to what was expected: \nExpected: %+v \nResult: %+v", expectedResult, result)
			}
		})
	}
}
