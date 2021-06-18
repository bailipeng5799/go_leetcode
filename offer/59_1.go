package main
//每次遍历窗口中的值得到最大值
//func maxSlidingWindow(nums []int, k int) []int {
//	res := []int{}
//	if len(nums) == 0{
//		return res
//	}
//	for i := 0;i <= len(nums) - k; i++{
//		max := nums[i]
//		for j := i + 1;j < i + k;j++{
//			if max < nums[j]{
//				max = nums[j]
//
//			}
//		}
//		res = append(res,max)
//	}
//	return res
//}
func maxSlidingWindow(nums []int,k int) []int {
	n := len(nums)
	if n == 0 || k == 0{
		return []int{}
	}
	res := make([]int,0)
	deque := make([]int,n - k + 1)
	//初始化第一个窗口
	for i := 0;i < k ;i++{
		//先添加进去一个元素之后
		//判断每一个元素队列中元素的大小如果nums[i] > 列尾元素
		for len(deque) > 0 && deque[len(deque) - 1] < nums[i]{
			deque = deque[:len(deque)-1]
		}
		deque = append(deque,nums[i])
	}
	res = append(res,deque[0])
	for i := k ;i < n ;i++{
		//当离开滑动窗口的值恰好是滑动窗口要删除的元素
		if deque[0] == nums[i-k]{
			deque = deque[1:]
		}
		//每次得到这个窗口的最大值，碰到比队列中元素大的直接删除队列中元素
		for len(deque) > 0 && deque[len(deque) - 1] < nums[i]{
			deque = deque[:len(deque)-1]
		}
		//将刚刚大元素进行添加
		deque= append(deque, nums[i])
		res = append(res,deque[0])

	}
	return res
}
