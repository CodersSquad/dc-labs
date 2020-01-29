package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	t := map[string]int{}
	p := strings.Split(s, " ")
	for _,word := range p{
		t[word]++
	}
	
	
	return t
	
}

func main() {
	wc.Test(WordCount)
}
