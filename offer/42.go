package main
//简单动态规划
//状态转移方程
//dp[i] = max(nums[i],dp[i-1]+nums[i])
func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i],nums[i]+nums[i-1])
		if nums[i] > nums[0] {
			nums[0] = nums[i]
		}
	}
	return nums[0]
}
func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}
