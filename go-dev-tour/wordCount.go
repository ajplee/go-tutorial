package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordArray := strings.Fields(s)

	histogram := make(map[string]int)
	for i := 0; i < len(wordArray); i++ {
		_, ok := histogram[wordArray[i]]
		if ok {
			histogram[wordArray[i]]++
		} else {
			histogram[wordArray[i]] = 1
		}
	}
	return histogram
}

func main() {
	wc.Test(WordCount)
}
