package main
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	//先对nums中的个数进行判断，如果小于2直接返回false
	//使用简单的滑动窗口算法，left,right分别为左指针和右指针
	//因为i,j不能重复所以 right = left + 1
	//由于题目需要满足abs（i - j） <= k,所以我们只需要right - left <= k,窗口大小为k
	//因为right一直在left的右边，所以将此条件设置为内层循环的条件
	//外层循环left不能大于nums长度，内存循环right不能大于nums长度并且right - left <= k
	//如果满足abs（nums[left] - nums[right]) <= t 直接返回

	if len(nums) < 2{
		return false
	}
	left,right := 0,1
	for left < len(nums){
		right = left + 1
		for right < len(nums) && right - left <= k{
			if abs(nums[left] - nums[right]) <= t{
				return true
			}
			right++
		}
		left++
	}
	return false
}
func abs(a int)int{
	if a > 0{
		return a
	}
	return -a
}
