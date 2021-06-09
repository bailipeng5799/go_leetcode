package main
//排序数组中的搜索问题，首先想到 二分法 解决。
func missingNumber2(nums []int) int {
	left, right := 0, len(nums) - 1
	for left <= right{
		mid := (left + right) / 2
		if nums[mid] == mid{
			left = mid + 1
		}else{
			right = mid - 1 // 如果去掉-1那么循环将不会停止
		}
	}
	return left
	//如果是下面这种写法那么将漏掉[0]输出1这种情况
	//left, right := 0, len(nums) - 1
	//for left < right{
	//	mid := (left + right) / 2
	//	if nums[mid] == mid{
	//		left = mid + 1
	//	}else{
	//		right = mid
	//	}
	//}
	//return left
}
func missingNumber(nums []int) int {
	for index,value := range nums{
		if index != value{
			return index
		}
	}
	return len(nums)
}