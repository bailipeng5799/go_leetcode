package main
//450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val { //进入左子树
		root.Left = deleteNode(root.Left,key)
		return root
	}
	if key > root.Val { //进入右子树
		root.Right = deleteNode(root.Right,key)
		return root
	}
	if root.Right == nil{ //到此说明找到了要删除的值，root为要删除的值root.Right为空直接返回left同理
		return root.Left
	}
	if root.Left == nil{
		return root.Right
	}
	minNode := root.Right //存在左右都不为空，先找到要删除的值的右子树的左下角的值也就是右子树中最小值
	//也就是中序遍历的后继节点
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	root.Val = minNode.Val //将此值赋给root
	root.Right = DeleteMinNode(root.Right)//删除此节点
	return root
}
func DeleteMinNode(root * TreeNode)*TreeNode{
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = DeleteMinNode(root.Left)
	return root
}