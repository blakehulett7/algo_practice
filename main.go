package main

import (
	"fmt"
	"unicode"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		num := nums[mid]

		if num < target {
			left = mid + 1
			continue
		}

		if num != target {
			right = mid - 1
			continue
		}

		return mid
	}

	return -1
}
