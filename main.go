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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	to_visit := []*TreeNode{root}
	result := [][]int{}

	for len(to_visit) != 0 {
		list_size := len(to_visit)
		list := []int{}
		for i := 0; i < list_size; i++ {
			current := to_visit[0]
			to_visit = to_visit[1:]
			if current == nil {
				continue
			}
			to_visit = append(to_visit, current.Left, current.Right)
			list = append(list, current.Val)
		}

		if len(list) == 0 {
			break
		}
		result = append(result, list)
	}

	return result
}
