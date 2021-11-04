package main
//连续子数组的最大和
//dp
func FindGreatestSumOfSubArray( array []int ) int {
	// write code here
	res := array[0]
	for i := 1; i < len(array); i++ {
		array[i] += max(array[i-1],0)
		res = max(res,array[i])
	}
	return res
	//dp := make([]int,len(array))
	//dp[0] = array[0]
	//maxInt := dp[0]
	//for i := 1; i < len(array); i++ {
	//	dp[i] = max(array[i],dp[i-1]+array[i])
	//	if dp[i] > maxInt{
	//		maxInt = dp[i]
	//	}
	//}
	//return maxInt
}
func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}
