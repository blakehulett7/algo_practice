package main

import (
	"fmt"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}

	for _, num := range nums {
		counts[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range counts {
		buckets[count] = append(buckets[count], num)
	}

	i := 0
	result := []int{}
	for j := len(buckets) - 1; i < k; j-- {
		bucket := buckets[j]
		for _, num := range bucket {
			result = append(result, num)
			i++
		}
	}

	return result
}
