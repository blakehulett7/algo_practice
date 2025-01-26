package main

import (
	"fmt"
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

	srunes := []rune(s)
	scounts := [26]int{}

	trunes := []rune(t)
	tcounts := [26]int{}

	for i := 0; i < len(s); i++ {
		scounts[srunes[i]-'a']++
		tcounts[trunes[i]-'a']++
	}

	return scounts == tcounts
}
