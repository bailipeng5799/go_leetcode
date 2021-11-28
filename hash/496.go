package main
//496. 下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int,len(nums1))
	for i := range res {    //将res初始化为-1
		res[i] = -1
	}
	mp := map[int]int{}
	for i, v := range nums1 { //将nums1存入map中
		mp[v] = i
	}
	//单调递减的一个栈，栈顶元素最小，如若遍历元素比栈顶元素大那么在map中查找栈顶元素并且输出
	stack := []int{}
	stack = append(stack,0) //stack中存放nums2的下标
	for i := 1; i < len(nums2); i++ { //i从1开始进行比较，如果nums2中下一个元素比栈中最后一个元素大，
		//查询该元素是否在map中存在，存在就的到下标
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			if _,ok := mp[nums2[top]]; ok {
				index := mp[nums2[top]]
				res[index] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)
	}
	return res
}

