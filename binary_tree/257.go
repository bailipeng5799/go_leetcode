package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//257. 二叉树的所有路径
//本题我使用了深度优先遍历的算法
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var DFS func(root *TreeNode,s string)
	DFS = func (root *TreeNode,s string){
		//判断为叶子节点时将root.Val拼接进这个路径中并返回
		if root.Left == nil && root.Right == nil {
			v := s + strconv.Itoa(root.Val)
			res = append(res,v)
			return
		}
		//如果不是叶子节点那么将此节点值拼接
		s = s + strconv.Itoa(root.Val) + "->"
		if root.Left != nil {
			DFS(root.Left,s)
		}
		if root.Right != nil {
			DFS(root.Right,s)
		}
	}
	DFS(root,"")
	return res
}

