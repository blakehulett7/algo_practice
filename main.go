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

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}

func isSubtree(root, subroot *TreeNode) bool {
	if subroot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if isSameTree(root, subroot) {
		return true
	}

	return isSubtree(root.Left, subroot) || isSubtree(root.Right, subroot)
}
