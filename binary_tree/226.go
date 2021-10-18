package main
//翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	//交换的代码 先进入最底层在交换
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}
//使用栈实现
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	stack = append(stack,root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
		if node.Left != nil {
			stack = append(stack,node.Left)
		}
		if node.Right != nil {
			stack = append(stack,node.Right)
		}
	}
	return root
}