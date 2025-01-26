package main

import (
	"fmt"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

func twoSum(nums []int, target int) []int {
	idx_map := map[int]int{}

	for idx, num := range nums {
		diff := target - num
		prev_idx, seen := idx_map[diff]
		if !seen {
			idx_map[num] = idx
			continue
		}

		return []int{prev_idx, idx}
	}

	return []int{}
}
