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

type Stack []rune

func (s *Stack) push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) pop() rune {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) peek() rune {
	return s[len(s)-1]
}

func isValid(s string) bool {
	stack := Stack{}
	valid_map := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	runes := []rune(s)

	for _, r := range runes {
		opener, is_closer := valid_map[r]
		if !is_closer {
			stack.push(r)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if opener != stack.pop() {
			return false
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
