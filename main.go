package main

import (
	"fmt"
	"strings"
)

func main() {
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

func scratch(tokens []string) int {
	stack := Stack{}

	for _, token := range tokens {
		if len(stack) == 0 {
			stack.push(token)
			continue
		}

		if !strings.Contains("+-*/", token) {
			stack.push(token)
			continue
		}

		num_1, num_2 := stack.pop(), stack.pop()
		stack.push(evaluate(token, num_1, num_2))
	}
}

func evaluate(token string, num_1, num_2 int) int {
	switch token {
	case "+":
		return num_1 + num_2
	case "-":
		return num_1 - num_2
	case "*":
		return num_1 * num_2
	case "/":
		return num_1 / num_2
	default:
		return 0
	}
}
