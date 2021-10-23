package main
import "math"
//530. 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	var res []int
	findMin(root,&res)
	min := math.MaxInt64
	for i := 1; i < len(res); i++ {
		tmp := res[i] - res[i-1]
		if tmp < min {
			min = tmp
		}
	}
	return min
}
func findMin(root *TreeNode,res *[]int){
	if root == nil{
		return
	}
	findMin(root.Left,res)
	*res = append(*res,root.Val)
	findMin(root.Right,res)
}
