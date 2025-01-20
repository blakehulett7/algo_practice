package main

import (
	"fmt"
	"slices"
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

	s_count := make([]int, 26)
	t_count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		s_count[s[i]-'a']++
		t_count[t[i]-'a']++
	}

	return slices.Equal(s_count, t_count)
}
