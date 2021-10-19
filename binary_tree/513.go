package main

import "math"

//513.找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue,root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			//i == 0 说明是最左边的元素
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
		}
	}
	return res
}
//513.找树左下角的值
//递归
func findBottomLeftValue2(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	maxDeep, res := math.MinInt32, 0
	var findleftval func(root *TreeNode,deep int)
	findleftval = func(root *TreeNode,deep int) {
		//因为第一个进入的元素一定是最左边的所以不用进行判定最左 而是 判断叶子节点的深度
		if root.Left == nil && root.Right == nil {
			if deep > maxDeep {
				res = root.Val
				maxDeep = deep
			}
		}
		if root.Left != nil {
			deep++
			findleftval(root.Left,deep)
			deep--//回溯
		}
		if root.Right != nil {
			deep++
			findleftval(root.Right,deep)
			deep--//回溯
		}
		return
	}
	findleftval(root,maxDeep)
	return res

}