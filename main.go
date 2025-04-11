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

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	scounts := [26]int{}
	tcounts := [26]int{}

	for i := 0; i < len(s); i++ {
		scounts[s[i]-'a']++
		tcounts[t[i]-'a']++
	}

	return scounts == tcounts
}
