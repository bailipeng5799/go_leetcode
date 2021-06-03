package main

import "strconv"
//dp类似于爬楼梯
func translateNum(num int) int {
	if num < 10{
		return 1
	}
	src := strconv.Itoa(num)
	dp := make([]int,len(src))
	dp[0] = 1
	dp[1] = 1
	if src[:2] >= "10" && src[:2] <= "25"{
		dp[1] = 2
	}
	for i := 2;i < len(src); i++{
		if src[i-1:i+1] >= "10" && src[i-1:i+1] <= "25"{
			dp[i]=dp[i-1]+dp[i-2]
		}else{
			dp[i] = dp[i-1]
		}
	}
	return dp[len(src)-1]
	// p, q, r := 0, 0, 1
	// for i := 0; i < len(src); i++ {
	//     p, q, r = q, r, 0
	//     r += q
	//     if i == 0 {
	//         continue
	//     }
	//     pre := src[i-1:i+1]
	//     if pre <= "25" && pre >= "10" {
	//         r += p
	//     }
	// }
	// return r
}
