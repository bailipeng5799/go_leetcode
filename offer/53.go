package main
func search(nums []int, target int) int {
	if len(nums) == 0{
		return 0
	}
	left := helper(nums,0,len(nums) - 1, target)//先二分拿到比target的第一个下标
	right := helper(nums,left,len(nums) - 1,target + 1)//拿到比target+1的第一个下标
	return right - left
}
func helper(nums []int, left, right, target int) int {
	for left <= right {//必须要等于因为如果只存在一个元素那么将不会进入循环在拿到target+1的下标时会直接返回
		mid := (left + right)/2
		if nums[mid] < target{
			left = left + 1
		}else{
			right = right - 1
		}
	}
	return left
}
//每次都拿到target的最后一个元素的下标+1
func helper2(nums []int, target int) int {
	left,right := 0,len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return left
}


