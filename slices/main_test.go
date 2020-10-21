package main

import (
	"fmt"
	"testing"
)

func Test_Pic(t *testing.T) {
	cases := []struct {
		nLines   int
		nColumns int
	}{
		{nLines: 2, nColumns: 2},
		{nLines: 1, nColumns: 5},
		{nLines: 5, nColumns: 5},
		{nLines: 15, nColumns: 30},
		{nLines: 50, nColumns: 100},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Test N%d", i), func(t *testing.T) {
			t.Parallel()

			picture := Pic(c.nLines, c.nColumns)
			if len(picture) != c.nLines {
				t.Errorf("The number of lines in the picture do not match. \n Expected: %d \n Got: %d", c.nLines, len(picture))
			}

			for _, line := range picture {
				if len(line) != c.nColumns {
					t.Errorf("One of the lines does not have the expected number of columns. \n Expected: %d \n Got: %d", c.nColumns, len(line))
				}
			}
		})
	}
}
