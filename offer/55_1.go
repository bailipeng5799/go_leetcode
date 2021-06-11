package main
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	depth := max(maxDepth(root.Left),maxDepth(root.Right)) + 1
	return depth
}
func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}
