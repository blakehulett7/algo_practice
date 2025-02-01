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
	node := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
			node = node.Next
			continue
		}

		node.Next = list2
		list2 = list2.Next
		node = node.Next
	}

	if list2 == nil {
		node.Next = list1
		return dummy.Next
	}

	node.Next = list2
	return dummy.Next
}
