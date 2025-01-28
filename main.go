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

type Stack []rune

func (s *Stack) push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) pop() rune {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func isValid(s string) bool {
	stack := Stack{}
	valid := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	runes := []rune(s)

	for _, r := range runes {
		open, is_closer := valid[r]
		if !is_closer {
			stack.push(r)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack.pop() != open {
			return false
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
