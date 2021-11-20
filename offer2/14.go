package main

func cuttingRope(n int) int {
	dp := make([]int,n+1)
	dp[2] = 1 //长度为2的绳子最大乘积
	for i := 3; i <= n; i++ { //从3开始
		for j := 2;j < i; j++ { //代表剪断的部分最少为2
			dp[i] = max(dp[i],max(j*(i-j),j*dp[i-j]))
			//j*(i-j) 为剪断两部分的乘积  j*dp[i-j]  剪断的部分*之前所剪断部分的dp
		}
	}
	return dp[n]
}
func max(a,b int)int {
	if a > b {
		return a
	}
	return b
}
