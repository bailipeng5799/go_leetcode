package main

import "fmt"

func numWays(n int) int {
	if n == 1 || n == 0{
		return 1
	}
	if n==2{
		return 2
	}
	dp := make([]int,n + 1)
	dp[1] = 1
	dp[2] = 2
	//方程 dp[i]=dp[i-1]+dp[i-2]
	for i := 3;i <= n ;i++{
		dp[i] = (dp[i-1]+dp[i-2])%1000000007
	}
	return dp[n]
}
func numWays2(n int) int {
	f0,f1 := 1,2
	if n == 1 || n == 0{
		return 1
	}
	if n==2{
		return 2
	}
	for i := 3; i <= n; i++{
		f0,f1 = f1,(f1+f0)%1000000007
	}
	return f1

}
func main(){
	result := numWays2(7)
	fmt.Println(result)
}
