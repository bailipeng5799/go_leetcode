package main
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	//将四个数组看成两个两个一组
	//先将前两个数组中可能存在的和的情况全部存入map中
	//然后判断后两个数组的和的相反数是否存在如果存在那么count++
	countAB := make(map[int]int)
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			countAB[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += countAB[0-(v3+v4)]
		}
	}
	return count
}
