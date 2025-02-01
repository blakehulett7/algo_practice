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

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]bool{}
	current := head

	for current != nil {
		if seen[current] {
			return true
		}

		seen[current] = true
		current = current.Next
	}

	return false
}
