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

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	runes := []rune(s)

	for left < right {
		left_char := unicode.ToLower(runes[left])
		right_char := unicode.ToLower(runes[right])

		if !unicode.IsDigit(left_char) && !unicode.IsLetter(left_char) {
			left++
			continue
		}

		if !unicode.IsDigit(right_char) && !unicode.IsLetter(right_char) {
			right--
			continue
		}

		if left_char != right_char {
			return false
		}

		left++
		right--
	}

	return true
}
