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
	fmt.Println(WordCount("one fish, two fish, red fish, blue fish"))
}
