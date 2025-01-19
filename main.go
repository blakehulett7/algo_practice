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

func (s *Stack) push(num int) {
	*s = append(*s, num)
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
		switch token {
		case "+":
			num_2 := stack.pop()
			num_1 := stack.pop()
			stack.push(num_1 + num_2)
		case "-":
			num_2 := stack.pop()
			num_1 := stack.pop()
			stack.push(num_1 - num_2)
		case "*":
			num_2 := stack.pop()
			num_1 := stack.pop()
			stack.push(num_1 * num_2)
		case "/":
			num_2 := stack.pop()
			num_1 := stack.pop()
			stack.push(num_1 / num_2)
		default:
			num, _ := strconv.Atoi(token)
			stack.push(num)
		}
	}
	return stack.pop()
}
