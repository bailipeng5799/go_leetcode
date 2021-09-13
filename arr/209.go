package main
func minSubArrayLen(target int, nums []int) int {
	//滑动窗口
	left, right := 0, 0
	length := len(nums)
	res := length + 1
	sum := 0
	for right < length {//当右指针移动到末尾后如果左指针右移时没有满足条件那么后续就不会存在
		sum += nums[right]
		for sum >= target { //这时右指针不动将左指针右移
			tmp := right - left + 1 //临时存储结果
			if tmp < res { //获取最小值
				res = tmp
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if res == length + 1{ //说明res没有改变也就是没有满足条件的字数组
		return 0
	}
	return res
}

