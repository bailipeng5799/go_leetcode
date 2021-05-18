package main

func exchange(nums []int) []int {
	result1 := []int{}
	result2 := []int{}
	if len(nums) == 0{
		return result1
	}
	for i:=0;i<len(nums);i++{
		if nums[i]%2==1{
			result1 = append(result1,nums[i])
		}else{
			result2 = append(result2,nums[i])
		}
	}
	return append(result1,result2...)
}
//双指针left从前遍历，找到偶数，right从右遍历找到奇数，二者交换继续遍历直至left==right
func exchange2(nums []int) []int {
	left,right := 0,len(nums)-1
	for left < right{
		if nums[left] & 1 == 1{
			left++
			continue
		}
		if nums[right] & 1 == 0{
			right--
			continue
		}
		nums[left],nums[right] = nums[right],nums[left]
		left++
		right--
	}
	return nums
}
