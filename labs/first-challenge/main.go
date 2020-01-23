package main

import (
	"fmt"
)

func main() {
	calculate("abcabcbb")
	calculate("bbbbbb")
	calculate("pwwkew")
}

type option struct {
	prefix string
	set    map[string]bool
	len    int
}

func calculate(data string) {
	options := map[string]option{}
	listOfP := []string{}
	//Create Set
	prefix := ""
	set := map[string]bool{}
	for i, e := range data {
		curLen := len(set)
		set[string(e)] = true

		if len(set) == curLen {
			isIn := false
			for _, e := range listOfP {
				if prefix == e {
					isIn = true
				}
			}
			if !isIn {
				options[prefix] = option{prefix: prefix, set: set, len: len(set)}
				listOfP = append(listOfP, prefix)
			}

			set = map[string]bool{}
			set[string(e)] = true
			prefix = ""
		}

		prefix += string(e)
		if i == len(data)-1 {
			options[prefix] = option{prefix: prefix, set: set, len: len(set)}
			set = map[string]bool{}
		}
	}

	//Filter
	maxN := 0
	for _, value := range options {
		if value.len >= maxN {
			maxN = value.len
		}
	}

	fmt.Println(maxN)
}
