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

type Stack [][]int

func (s *Stack) push(num []int) {
	*s = append(*s, num)
}

func (s *Stack) pop() []int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) peek() []int {
	return s[len(s)-1]
}

func largestRectangleArea(heights []int) int {
	largest := 0
	stack := Stack{}

	idx := 0
	for idx < len(heights) {
		height := heights[idx]
		if len(stack) == 0 {
			stack.push([]int{idx, height})
			idx++
			continue
		}
		prev_height := stack.peek()[1]
		if prev_height <= height {
			stack.push([]int{idx, height})
			idx++
			continue
		}

		prev_idx := stack.pop()[0]
		width := idx - prev_idx
		area := prev_height * width

		if largest < area {
			largest = area
		}
	}

	for len(stack) != 0 {
		prev_idx := stack.peek()[0]
		width := idx - prev_idx
		prev_height := stack.pop()[1]
		area := prev_height * width
		if largest < area {
			largest = area
		}
	}

	return largest
}
