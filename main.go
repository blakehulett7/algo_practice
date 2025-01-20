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

type Stack []float64

func (s *Stack) push(num float64) {
	*s = append(*s, num)
}

func (s *Stack) pop() float64 {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) peek() float64 {
	return s[len(s)-1]
}

// Need to sort this by position first
func carFleet(target int, position, speed []int) int {
	stack := Stack{}

	cars := [][2]int{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, [2]int{position[i], speed[i]})
	}

	slices.SortFunc(cars, func(a, b [2]int) int {
		return b[0] - a[0]
	})

	for _, car := range cars {
		time := float64(target-car[0]) / float64(car[1])

		if len(stack) == 0 {
			stack.push(time)
			continue
		}

		prev_time := stack.peek()
		if time > prev_time {
			stack.push(time)
		}
	}

	return len(stack)
}
