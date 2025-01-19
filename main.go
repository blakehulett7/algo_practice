package main

import (
	"fmt"
	"strconv"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

type Stack []int

func (s *Stack) push(element int) {
	*s = append(*s, element)
}

func (s *Stack) pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) peek() int {
	return s[len(s)-1]
}

func scratch(tokens []string) int {
	stack := Stack{}

	for _, token := range tokens {
		fmt.Printf("stack: %s, token: %s", stack, token)
		switch token {
		case "+":
			num_1 := stack.pop()
			num_2 := stack.pop()
			stack.push(num_1 + num_2)
		case "-":
			num_1 := stack.pop()
			num_2 := stack.pop()
			stack.push(num_1 - num_2)
		case "*":
			num_1 := stack.pop()
			num_2 := stack.pop()
			stack.push(num_1 * num_2)
		case "/":
			num_1 := stack.pop()
			num_2 := stack.pop()
			stack.push(num_2 / num_1)
		default:
			num, _ := strconv.Atoi(token)
			stack.push(num)
		}
	}
	return stack.pop()
}
