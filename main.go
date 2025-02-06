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

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{1, 0}
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if left[0] != 1 || right[0] != 1 {
			return []int{0, 0}
		}

		if abs(right[1]-left[1]) > 1 {
			return []int{0, 0}
		}

		return []int{1, max(left[1], right[1]) + 1}
	}

	is_balanced := dfs(root)[0]

	if is_balanced != 1 {
		return false
	}

	return true
}
