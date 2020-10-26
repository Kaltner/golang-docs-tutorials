package main

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := map[float64]float64{
		1:      1,
		1.5:    1.22,
		2:      1.41,
		3:      1.73,
		4:      2,
		5:      2.24,
		120:    10.95,
		-1:     0,
		-2:     0,
		-100:   0,
		-12340: 0,
	}

	for parameter, expectedResult := range cases {
		parameter, expectedResult := parameter, expectedResult
		t.Run(fmt.Sprintf("Test N%f", parameter), func(t *testing.T) {
			// t.Parallel()

			result, err := Sqrt(parameter)
			if parameter < 0 {
				if err == nil {
					t.Errorf("Squareroot is accepting negative numbers: %+v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Error converting number: %+v", err)
				}
				if result != expectedResult {
					t.Errorf("Result is not equal to what was expected: \nExpected: %+v \nResult: %+v", expectedResult, result)
				}
			}
		})
	}
}
