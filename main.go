package main

import "fmt"

func main() {

}

func fmtin() {
	fmt.Println()
}

func scratch(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_runes := []rune(s)
	t_runes := []rune(t)

	s_hash := map[rune]int{}
	t_hash := map[rune]int{}

	for i := 0; i < len(s); i++ {
		_, exists := s_hash[s_runes[i]]
		if !exists {
			s_hash[s_runes[i]] = 0
		}
		s_hash[s_runes[i]]++

		_, exists = t_hash[s_runes[i]]
		if !exists {
			t_hash[t_runes[i]] = 0
		}
		t_hash[t_runes[i]]++
	}

	for key, s_value := range s_hash {
		t_value := t_hash[key]
		if s_value != t_value {
			return false
		}
	}

	return true
}
