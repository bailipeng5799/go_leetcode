package main
//平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//判断如果左子树和右子树中有一个返回为false那么直接返回false
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	//判断左子树和右子树最大深度的绝对值是否大于1如果大于1 返回false
	if abs(DFS(root.Left)-DFS(root.Right)) > 1{
		return false
	}
	return true

}
func DFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(DFS(root.Left),DFS(root.Right))+1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int{
	if a < 0 {
		return -a
	}
	return a
}
