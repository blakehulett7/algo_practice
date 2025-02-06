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

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
			continue
		}

		current.Next = list2
		current = current.Next
		list2 = list2.Next
	}

	if list2 == nil {
		current.Next = list1
		return dummy.Next
	}

	current.Next = list2
	return dummy.Next
}
