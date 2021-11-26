package main
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//因为root.Left会被置换
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root

}
