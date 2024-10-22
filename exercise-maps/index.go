package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	slice := strings.Split(s, " ")
	res := map[string]int{}
	for _, x := range slice {
		res[x]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
