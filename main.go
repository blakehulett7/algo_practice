package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

func is_anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	scount := [26]int{}
	tcount := [26]int{}

	for i := 1; i < len(s); i++ {
		schar := s[i]
		tchar := t[i]

		scount[schar-'a']++
		tcount[tchar-'a']++
	}

	return scount == tcount
}
