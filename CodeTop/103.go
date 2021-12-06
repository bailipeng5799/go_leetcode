package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil{
		return nil
	}
	level := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp,node.Val)
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
		}
		if level % 2 == 1 {
			for j := 0; j < len(tmp)/2; j++ {
				tmp[j], tmp[len(tmp)-j-1] = tmp[len(tmp)-j-1], tmp[j]
			}
		}
		result = append(result,tmp)
		level++
	}
	return result
}

