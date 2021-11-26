package main
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left,root.Right)
}
func compare(a,b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return compare(a.Left,b.Right) && compare(a.Right,b.Left)
}