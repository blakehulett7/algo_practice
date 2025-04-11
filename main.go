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

type Stack []int

func (s *Stack) push(val int) {
	*s = append(*s, val)
}

func (s Stack) peek() int {
	return s[len(s)-1]
}

func (s *Stack) pop() int {
	top := s.peek()
	*s = (*s)[:len(*s)-1]
	return top
}
