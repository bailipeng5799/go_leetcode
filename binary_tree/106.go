package main
//106. 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	nodeval := postorder[len(postorder)-1]
	left := findRootIndex(inorder,nodeval)//确定左子树的范围
	root := &TreeNode{
		Val :nodeval,
		Left: buildTree(inorder[:left],postorder[:left]),
		Right:buildTree(inorder[left+1:],postorder[left:len(postorder)-1]),
	}
	return root
}
func findRootIndex(inorder []int,target int) (index int){
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}