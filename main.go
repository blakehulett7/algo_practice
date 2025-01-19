package main

import "fmt"

func main() {
	scratch("")
}

func fmtin() {
	fmt.Println()
}

type Stack []any

func (s *Stack) push(element any) {
	*s = append(*s, element)
}

func (s *Stack) pop() any {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) peek() any {
	return s[len(s)-1]
}

func scratch(str string) bool {
	valid_map := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	runes := []rune(str)
	stack := Stack{}

	for _, r := range runes {
		if len(stack) == 0 {
			stack.push(r)
			continue
		}

		left, is_right := valid_map[r]
		if !is_right {
			stack.push(r)
			continue
		}

		if stack.pop() != left {
			return false
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
