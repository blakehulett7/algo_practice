package main

import "fmt"

func main() {

}

func fmtin() {
	fmt.Println()
}

func scratch(strs []string) [][]string {
	anagram_map := map[[26]int][]string{}

	for _, str := range strs {
		key := [26]int{}
		str_rune := []rune(str)
		for _, r := range str_rune {
			key[r-'a']++
		}

		_, seen := anagram_map[key]
		if !seen {
			anagram_map[key] = []string{}
		}
		anagram_map[key] = append(anagram_map[key], str)
	}

	result := [][]string{}
	for _, value := range anagram_map {
		result = append(result, value)
	}
	return result
}
