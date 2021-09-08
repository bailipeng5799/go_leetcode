package main

func maxProfit2(prices []int) int {
	//相当于求每一次上涨的空间大小，如果tmp > 0 代表上涨了，
	//如果< 0说明下降了那么就不将tmp加入res中
	//所以最终求到的res值时每一天上升结果的总和
	res := 0
	for i := 1; i < len(prices); i++{
		tmp := prices[i] - prices[i-1]
		if tmp > 0{
			res += tmp
		}
	}
	return res
}
//func maxProfit2(prices []int) int {
//	n := len(prices)
//	if n < 2 {
//		return 0
//	}
//	dp := make([][2]int,n)
//	//dp[i][0] 表示第 i 天交易完后手里有现金
//	//dp[i][1] 表示第 i 天交易完后手里有股票
//	dp[0][0] = 0
//	dp[0][1] = -prices[0]
//	for i := 1; i < n; i++{
//		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
//		//第i天的现金最大值为前一天的现金最大值，或者前一天买的股票今天卖出一共的现金
//		dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])
//		//第i天的股票等于前一天的股票，或者前一天有现金今天买进的股票  其实就是此处得到的就是净利润
//	}
//	return dp[n-1][0]
//}
//
//func max(a,b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}