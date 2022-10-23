package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	var r = make(map[string]int)
	split := strings.Fields(s)
	for _, v := range split {
		r[v]++
	}

	return r
}

func main() {
	wc.Test(WordCount)
}
