package main

import "testing"

func testSqrt(t *testing.T) {
	cases := map[float64]int{
		1: 1,
		2: 1,
		3: 1,
		4: 2,
		5: 2,
	}

	for number, expectedResult := range cases {
		number, expectedResult := number, expectedResult
		t.Run(string(number), func(t *testing.T) {
			t.Parallel()

			result := Sqrt(number)

		})
	}
}
