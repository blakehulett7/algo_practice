package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

type ListNode Struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var previous *ListNode
    current := head

    for current != nil {
        next := current.Next
        current.Next = previous
        previous = current
        current = next
    }

    return previous
}
