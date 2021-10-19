package main
//112. 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0
	flag := false
	if root == nil {
		return flag
	}
	var findTargetSum func(root *TreeNode)
	findTargetSum = func(root *TreeNode) {
		sum += root.Val
		//先判断叶子节点是否为叶子节点   判断sum是否等于targetsum
		if root.Left == nil && root.Right == nil {
			if sum == targetSum {
				flag = true
				return
			}
		}
		if root.Left != nil && !flag { //如果已经成功就不需要在进行递归
			findTargetSum(root.Left)
		}
		if root.Right != nil && !flag {
			findTargetSum(root.Right)
		}
		sum -= root.Val //回溯

	}
	findTargetSum(root)
	return flag
}