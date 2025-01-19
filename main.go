package main

import (
	"fmt"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

type Stack [][2]int

func (s *Stack) push(num_idx [2]int) {
	*s = append(*s, num_idx)
}

func (s *Stack) pop() [2]int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) peek() [2]int {
	return s[len(s)-1]
}

func dailyTemperatures(temperatures []int) []int {
	stack := Stack{}
	result := make([]int, len(temperatures))

	idx := 0
	for idx < len(temperatures) {
		temp := temperatures[idx]

		if len(stack) == 0 {
			stack.push([2]int{temp, idx})
			idx++
			continue
		}

		prev_temp := stack.peek()[0]
		if temp <= prev_temp {
			stack.push([2]int{temp, idx})
			idx++
			continue
		}

		prev_idx := stack.pop()[1]
		result[prev_idx] = idx - prev_idx
	}

	return result
}
