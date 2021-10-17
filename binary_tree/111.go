package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := math.MaxInt32
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil {
		min = Selectmin(min,minDepth(root.Left))
	}
	if root.Right != nil {
		min = Selectmin(min,minDepth(root.Right))
	}
	return min + 1
}
func Selectmin(a,b int)int {
	if a < b {
		return a
	}
	return b
}
