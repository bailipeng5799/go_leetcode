package main
func maxSlidingWindow(nums []int, k int) []int {

	queue := []int{}//此队列从队首到队尾由大到小
	push := func(i int) {//添加元素如果添加的元素比队尾的元素大删除队尾元素
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue,i) //注意操作的是元素下标
	}
	for i := 0; i < k; i++ { //先把前k个中最大的元素放到队首
		push(i)
	}
	n := len(nums)
	ans := make([]int,1,n-k+1)//这种定义 切片的长度为1 但是预留的空间长度为n-k+1
	ans[0] = nums[queue[0]] //第一个值为队首元素的值也就是前k个元素的最大值
	for i := k; i < n; i++ {//从第k个元素开始依次进行push每次push的到的最大素结果为队首元素
		push(i)             //在下面进行判断之后直接将队首元素添加到ans切片中
		for queue[0] <= i - k { //如果队首元素已经不在滑动窗口中那么删除队首元素  循环进行判断 确保队首元素一定在窗口中
			queue = queue[1:]
		}
		ans = append(ans,nums[queue[0]])
	}
	return ans

}