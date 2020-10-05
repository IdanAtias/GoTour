package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount - count words in a string
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	var words []string = strings.Fields(s)
	for _, word := range words {
		if _, ok := result[word]; !ok {
			result[word] = 1
		} else {
			result[word]++
		}
	}
	return result
}

func runEx3() {
	wc.Test(WordCount)
}
