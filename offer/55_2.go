package main
func isBalanced(root *TreeNode)bool{
	if root == nil{
		return true
	}
	if maxDepth(root.Left) - maxDepth(root.Right) <= 1 && maxDepth(root.Left) - maxDepth(root.Right) >= -1  && isBalanced(root.Left) && isBalanced(root.Right){
		return true
	}
	return false
}
func maxDepth(root *TreeNode)int{
	if root == nil{
		return 0
	}
	return max(maxDepth(root.Left),maxDepth(root.Right)) + 1
}
func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}