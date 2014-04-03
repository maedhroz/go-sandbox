package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	counters := map[string]int{}

	for _, word := range strings.Fields(s) {
		_, exists := counters[word]

		if !exists {
			counters[word] = 0
		}

		counters[word]++
	}

	return counters
}

func main() {
	fmt.Println(WordCount("the quick brown fox jumped over the lazy dog"))
}
