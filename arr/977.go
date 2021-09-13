package main

func sortedSquares(nums []int) []int {
	//数组的最大值只能是最右边或者最左边
	//创建一个新的数组从右往左存放最大值元素
	length := len(nums)
	left, right := 0, length-1
	newNum := make([]int,length)
	for length > 0 { //注意这里是大于0因为存入元素的时候我们使用的是length-1所以length到1的时候
		//数组已经存满了
		if (nums[left] * nums[left]) > (nums[right] * nums[right]){
			newNum[length-1] = nums[left] * nums[left]
			left++

		}else{
			newNum[length-1] = nums[right] * nums[right]
			right--
		}
		length--
	}
	return newNum
}
