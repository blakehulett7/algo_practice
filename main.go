package main

import (
	"fmt"
	"slices"
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
