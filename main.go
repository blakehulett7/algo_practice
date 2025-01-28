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

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head

	for current != nil {
		node := current.Next
		current.Next = previous
		previous = current
		current = node
	}

	return previous
}
