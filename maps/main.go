package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	splitedStr := strings.Fields(s)

	totalWords := make(map[string]int)
	for _, splitStr := range splitedStr {
		splitStr = strings.ToLower(splitStr)
		_, ok := totalWords[splitStr]
		if !ok {
			totalWords[splitStr] = 0
		}

		totalWords[splitStr] += 1
	}
	return totalWords
}

func main() {
	fmt.Printf("%+v \n", WordCount("Bettlejuice! Bettlejuice! Bettlejuice!"))
}
