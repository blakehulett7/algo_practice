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

func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{1, 0}
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left[0] == 0 || right[0] == 0 {
			return []int{0, 0}
		}

		if abs(right[1]-left[1]) > 1 {
			return []int{0, 0}
		}

		return []int{1, max(left[1], right[1]) + 1}
	}

	is_balanced := dfs(root)
	if is_balanced[0] == 0 {
		return false
	}

	return true
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}
