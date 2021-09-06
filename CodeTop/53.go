package main
//最大连续子序和
//使用动态规划的方法直接进行操作
func maxSubArray(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	//dp := make([]int,len(nums))
	//dp[0] = nums[0]
	//result := nums[0]
	//for i := 1; i < len(nums); i++{
	//	//动态规划方程
	//	dp[i] = max(dp[i-1]+nums[i],nums[i])
	//	result = max(result,dp[i])
	//}
	//可以不用申请动态数组直接在原数组上进行修改
	//将每一次在此位置得到的最大值直接进行覆盖
	result := nums[0]
	for i := 1; i < len(nums); i++{
		nums[i] = max(nums[i-1]+nums[i],nums[i])
		result = max(result,nums[i])
	}
	return result

}
func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}