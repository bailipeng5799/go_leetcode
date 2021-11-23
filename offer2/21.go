package main
//剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left] % 2 == 1{
			left++
		}
		for left < right && nums[right] % 2 == 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
