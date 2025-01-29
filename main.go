package main

import (
	"fmt"
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

	s_runes := []rune(s)
	s_counts := [26]int{}

	t_runes := []rune(t)
	t_counts := [26]int{}

	for i := 0; i < len(s); i++ {
		s_counts[s_runes[i]-'a']++
		t_counts[t_runes[i]-'a']++
	}

	return s_counts == t_counts
}
