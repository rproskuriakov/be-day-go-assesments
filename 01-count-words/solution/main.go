package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func CountWords(text string) map[string]int {
	counts := make(map[string]int)
	for _, w := range strings.Fields(text) {
		counts[strings.ToLower(w)]++
	}
	return counts
}

func main() {
	c := CountWords("Go is fun and go is fast")
	for _, w := range slices.Sorted(maps.Keys(c)) {
		fmt.Printf("%s: %d\n", w, c[w])
	}
}
