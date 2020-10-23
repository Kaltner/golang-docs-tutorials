package main

import (
	"fmt"
	"testing"
)

func Test_IPType_String(t *testing.T) {
	// randNums := []byte{}
	cases := []struct {
		ip             IPAddr
		expectedResult string
	}{
		{ip: IPAddr{127, 0, 0, 1}, expectedResult: "127.0.0.1"},
		{ip: IPAddr{8, 8, 8, 8}, expectedResult: "8.8.8.8"},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Test case N%d", i), func(t *testing.T) {
			t.Parallel()

			if c.ip.String() != c.expectedResult {
				t.Errorf("Error stringifying the IP Address. \n Expected: %s \n Got: %s", c.expectedResult, c.ip.String())
			}
		})
	}
}
