package main

import (
	"fmt"
	"strconv"
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
		switch token {
		case "+":
			num_1 := stack.pop().(int)
			num_2 := stack.pop().(int)
			stack.push(num_1 + num_2)
		case "-":
			num_1 := stack.pop().(int)
			num_2 := stack.pop().(int)
			stack.push(num_1 - num_2)
		case "*":
			num_1 := stack.pop().(int)
			num_2 := stack.pop().(int)
			stack.push(num_1 * num_2)
		case "/":
			num_1 := stack.pop().(int)
			num_2 := stack.pop().(int)
			stack.push(num_1 / num_2)
		default:
			num, _ := strconv.Atoi(token)
			stack.push(num)
		}
	}
	return stack.pop().(int)
}
