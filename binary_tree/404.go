package main
//404. 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	var findleft func(root *TreeNode)
	findleft = func(root *TreeNode) {
		//判断是否为左叶子节点
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			res += root.Left.Val
		}
		if root.Left != nil {
			findleft(root.Left)
		}
		if root.Right != nil {
			findleft(root.Right)
		}
	}
	findleft(root)
	return res

}

