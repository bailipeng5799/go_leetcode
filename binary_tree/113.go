package main
//113. 路径总和 II
func pathSum(root *TreeNode, targetSum int) [][]int {
	sum := 0
	// flag := false
	res := [][]int{}
	tmp := []int{}
	if root == nil {
		return nil
	}
	var findPath func(root *TreeNode)
	findPath = func(root *TreeNode) {
		sum += root.Val
		tmp = append(tmp,root.Val)
		if root.Left == nil && root.Right == nil && sum == targetSum{
			cp := make([]int, len(tmp)) //这里一定要copy 要么当你append之后直接添加tmp传输的是一个地址
			copy(cp, tmp)               //如果后续没有扩容操作那么tmp一直修改的是刚刚传入res中的值
			res = append(res,cp)
		}
		if root.Left != nil {
			findPath(root.Left)
		}
		if root.Right != nil {
			findPath(root.Right)
		}
		sum -= root.Val
		tmp = tmp[:len(tmp)-1]

	}
	findPath(root)
	return res
}