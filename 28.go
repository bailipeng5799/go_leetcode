package main
func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return compare(root.Left,root.Right)
}
func compare(A,B *TreeNode)bool{
	if A == nil && B == nil{
		return true
	}
	if A == nil || B == nil || A.Val != B.Val{
		return false
	}
	return compare(A.Left,B.Right) && compare(A.Right,B.Left)
}

