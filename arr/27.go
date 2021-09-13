package main
func removeElement(nums []int, val int) int {
	left,right := 0,0
	for right = 0;right<len(nums);right++{
		if val != nums[right]{//如果没有找到所要删除的元素left，right同时++，如果找到了right++，此后left和right相差一次，
			// 每次将覆盖前面的值，当right到最后后一个元素时left的值就是所要求的长度
			nums[left]=nums[right]
			left++
		}

	}
	return left//返回left因为每次交换完毕之后直接对left++使最后一次交换完后的left就等于所删除后数组的长度
	//可以使用没有一个需要删除的元素来进行测验
}
