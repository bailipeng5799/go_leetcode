package main

import "math"

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode //前一个指针
	var truncate func(node *TreeNode) bool
	truncate = func(node *TreeNode)bool{
		if node == nil {
			return true
		}
		leftRes := truncate(node.Left) //一直进入到树的最左下角的元素
		if prev != nil && node.Val <= prev.Val {
			//是左中右的顺序
			return false //这里的prev记录的是此node节点的上一个元素 如果他小于上一个元素那么直接返回false
		}
		prev = node //存储上一个元素
		rightRes := truncate(node.Right)
		return leftRes && rightRes
	}
	return truncate(root)
}
//此方法主要是基于回溯的思想是一个中序遍历
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil{
		//细节就是在循环内部进行添加root而不是在外部添加,因为在内部可以添加root.Right元素
		for root != nil { //每一次将root添加进来
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack =stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
