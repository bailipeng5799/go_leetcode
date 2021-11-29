package main
//二分
func search( nums []int ,  target int ) int {
	// write code here
	left, right := 0, len(nums)-1
	flag := false
	for left <= right { //必须要等于可能会存在查找到切片的最后一个元素的情况此时left和right相等
		mid := (left + right)/2
		if nums[mid] > target {
			right = mid-1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		for mid >= 0 && nums[mid] == target { //如果相等向前推进直到找到不是target的元素之后返回+1
			mid--
			flag = true
		}
		if flag {
			return mid+1
		}
	}
	return -1
}