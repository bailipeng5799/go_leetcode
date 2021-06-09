package main
func maxValue(grid [][]int) int {
	dp:=make([][]int,len(grid)+1)
	for i:=0;i<=len(grid);i++{
		dp[i] = make([]int,len(grid[0])+1)
	}
	// dp[i][j]表示从grid[0][0]到grid[i - 1][j - 1]时的最大价值
	//  多开一行dp用来遍历第一行元素
	for i:=1;i<len(dp);i++{
		for j:=1;j<len(dp[i]);j++{
			dp[i][j]=grid[i-1][j-1]+max(dp[i][j-1],dp[i-1][j])
		}
	}
	return dp[len(grid)][len(grid[0])]
}
func max(x,y int)int{
	if x<y{
		return y
	}
	return x
}
