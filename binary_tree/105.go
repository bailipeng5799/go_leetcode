package main
//105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	nodeval := preorder[0]
	left := findRootIndex(nodeval,inorder)//其实找到的是左子树的长度，找到了根节点的下标
	root := &TreeNode{
		Val : nodeval,
		//因为preorder为前序队列，第一个节点为根节点，所以从索引为1开始到left+1为左子树元素的个数
		Left:buildTree(preorder[1:left+1],inorder[:left]),//从0到left
		Right:buildTree(preorder[left+1:],inorder[left+1:]),//前序从left+1也就是右子树的第一个节点
	}
	return root
}
func findRootIndex(target int,inorder []int)(index int) {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}