package main

//type TreeNode struct{
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}
//前序遍历
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var Traversal func(root *TreeNode)
	Traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res,root.Val)
		Traversal(root.Left)
		Traversal(root.Right)
	}
	Traversal(root)
	return res
}

//前序遍历迭代
func preorderTraversal2(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	node := root
	//一直到栈中的元素个数为0 并且node == nil
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res,node.Val)
			stack = append(stack,node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
//中序遍历
func inorderTraversal(root *TreeNode) []int {
	var traversal func(root *TreeNode)
	res := []int{}
	traversal = func(root *TreeNode) {
		if root == nil{
			return
		}
		traversal(root.Left)
		res = append(res,root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

//中序遍历迭代
func inorderTraversal2(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack,node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res,node.Val)
		node = node.Right
	}
	return res
}
//后序遍历
func postorderTraversal(root *TreeNode) []int {
	var traversal func(root *TreeNode)
	res := []int{}
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		res = append(res,root.Val)
	}
	traversal(root)
	return res
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//后序遍历
func postorderTraversal2(root *TreeNode) []int {
	stack :=  []*TreeNode{}
	res := []int{}
	node := root
	var prev *TreeNode //用来存放前一个节点判断 右节点是否已经被存入结果中
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack,node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if  node.Right == nil || node.Right == prev {
			res = append(res,node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack,node)
			node = node.Right
		}
	}
	return res
}