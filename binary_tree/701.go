package main
//701. 二叉搜索树中的插入操作
//递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val:val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left,val)
	}else {
		root.Right = insertIntoBST(root.Right,val)
	}
	return root
}
//迭代
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val:val}
	}
	node := root
	var prev *TreeNode //用前置指针来存储上一个节点最终结束循环时val与上一个节点作比较
	for node != nil {   //判断是左节点还是右节点插入
		if val > node.Val {
			prev = node
			node = node.Right
		}else{      //特别注意不能使用if 如果使用if 可能导致先执行上面的node=node.Right 然后在执行这个
			prev = node
			node = node.Left
		}
	}
	if val > prev.Val {
		prev.Right = &TreeNode{Val:val}
	}
	if val < prev.Val {
		prev.Left = &TreeNode{Val:val}
	}
	return root
}