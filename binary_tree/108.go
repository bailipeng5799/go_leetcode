package main
//108. 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { //如果等于nil直接返回
		return nil
	}
	root := &TreeNode{  //取中间元素作为val值
		Val:nums[len(nums)/2],
		Left:nil,
		Right:nil,
	}
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2 + 1:])
	return root
}
