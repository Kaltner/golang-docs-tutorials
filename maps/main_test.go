package main

import (
	"fmt"
	"testing"
)

var loremIpsumCount = map[string]int{
	"vestibulum":  2,
	"vestibulum,": 1,
	"odio":        3,
	"ac.":         1,
	"ac":          2,
	"suscipit":    2,
	"neque":       2,
	"massa":       1,
	"massa,":      1,
	"malesuada":   1,
	"malesuada,":  1,
	"magna.":      1,
	"magna":       1,
	"ipsum":       2,
	"interdum":    1,
	"interdum,":   1,
	"id":          2,
	"etiam":       2,
	"dolor":       1,
	"dolor,":      1,
	"cursus":      2,
	"consectetur": 2,
	"a":           2,
	"vitae":       1,
	"vel":         1,
	"turpis,":     1,
	"tristique":   1,
	"tortor":      1,
	"tempus":      1,
	"sit":         1,
	"sem":         1,
	"sed":         1,
	"risus":       1,
	"purus":       1,
	"primis":      1,
	"pretium":     1,
	"praesent":    1,
	"porta,":      1,
	"placerat":    1,
	"nunc.":       1,
	"nisi":        1,
	"nec":         1,
	"maximus":     1,
	"mauris":      1,
	"lorem":       1,
	"libero,":     1,
	"lectus.":     1,
	"in":          1,
	"imperdiet.":  1,
	"feugiat":     1,
	"fermentum":   1,
	"faucibus.":   1,
	"fames":       1,
	"euismod":     1,
	"eu":          1,
	"et":          1,
	"eros,":       1,
	"enim.":       1,
	"elit.":       1,
	"elementum":   1,
	"efficitur":   1,
	"donec":       1,
	"dignissim":   1,
	"consequat":   1,
	"at,":         1,
	"ante":        1,
	"amet,":       1,
	"adipiscing":  1,
}

func Test_WordCount(t *testing.T) {
	cases := []struct {
		textInput      string
		expectedResult map[string]int
	}{
		{
			textInput: "Bettlejuice! Bettlejuice! Bettlejuice!",
			expectedResult: map[string]int{
				"Bettlejuice!": 3,
			},
		},
		{
			textInput:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec sem turpis, euismod a vestibulum at, consectetur nec magna. Mauris feugiat placerat libero, id cursus neque cursus ac. Interdum et malesuada fames ac ante ipsum primis in faucibus. Etiam vitae efficitur dolor, maximus tristique lectus. Praesent vestibulum, odio sed fermentum malesuada, purus magna consequat massa, a dignissim massa odio ac enim. Vestibulum porta, odio eu elementum interdum, neque tortor tempus eros, id suscipit nisi risus vel nunc. Etiam suscipit pretium imperdiet.",
			expectedResult: loremIpsumCount,
		},
	}

	for i, c := range cases {
		i, c = i, c
		t.Run(fmt.Sprintf("Test case N: %d", i), func(t *testing.T) {
			t.Parallel()

			wordCount := WordCount(c.textInput)

			if len(c.expectedResult) != len(wordCount) {
				t.Errorf("Number of words found different than the expected. \n Expected: %d \n Got: %d", len(c.expectedResult), len(wordCount))
			}

			for expectedStr, expectedNOcurrences := range c.expectedResult {
				count, ok := wordCount[expectedStr]
				if !ok {
					t.Errorf("String not found by the word count. \n %s", expectedStr)
				}

				if count != expectedNOcurrences {
					t.Errorf("Number of ocurrences of the %q word do not match \n Expected: %d \n Got: %d", expectedStr, expectedNOcurrences, count)
				}
			}
		})
	}
}
