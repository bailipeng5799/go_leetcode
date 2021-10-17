package main

import "math"
//515. 在每个树行中找最大值
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	list := []*TreeNode{}
	list = append(list,root)
	for len(list) > 0 {
		length := len(list)
		max := math.MinInt64
		for i := 0; i < length; i++ {
			node := list[0]
			list = list[1:]
			if node.Left != nil {
				list = append(list,node.Left)
			}
			if node.Right != nil {
				list = append(list,node.Right)
			}
			if node.Val > max {
				max = node.Val
			}
		}
		res = append(res,max)
	}
	return res
}
