package main
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}
