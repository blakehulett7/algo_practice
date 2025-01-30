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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max_length := 0

	var dfs func(*TreeNode) int
	dfs = func(current *TreeNode) int {
		if current == nil {
			return 0
		}

		left := dfs(current.Left)
		right := dfs(current.Right)

		max_length = max(max_length, left+right)
		return 1 + max(left, right)
	}

	dfs(root)
	return max_length
}
