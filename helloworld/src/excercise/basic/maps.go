package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	x := make(map[string]int)
	for _, char := range strings.Fields(s) {
		x[char] += 1
	}
	return x
}

func main() {
	wc.Test(WordCount)
}
