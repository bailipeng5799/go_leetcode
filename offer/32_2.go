package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	var res [][]int

	queue := []*TreeNode{root}
	for len(queue) != 0{
		tmp := []int{}
		i := len(queue)//每次循环这么多次是一层的长度
		for ;i>0;i--{
			node := queue[0]
			tmp = append(tmp,node.Val)
			queue = queue[1:]
			if node.Left != nil{
				queue = append(queue,node.Left)
			}
			if node.Right != nil{
				queue = append(queue,node.Right)
			}
		}
		res = append(res,tmp)
	}
	return res
}

