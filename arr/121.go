package main
func maxProfit(prices []int) int {
	//本题使用动态规划的解法
	dp := make([]int,len(prices))
	min := prices[0]//min用来存放股票的最低点
	for i := 1; i < len(prices);i++{
		if prices[i] < min{
			min = prices[i] //更新min的值
		}
		//动态转移方程，dp[i] 表示前 i 天的最大利润
		dp[i] = max(dp[i-1],prices[i]-min)
	}
	return dp[len(prices) - 1]
}
func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}
