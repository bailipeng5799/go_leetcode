package main

func intersection(nums1 []int, nums2 []int) []int {
	//先将元素放入一个map中
	flag := make(map[int]bool)
	for _, v := range nums1 {
		flag[v] = true
	}
	k := 0//用来存放交集的数量
	for _, v := range nums2{
		//如果map中存在这个元素将其置换在nums2中
		//将map中的值变为false将k++最后返回nums2[:k]切片中前k个元素
		if flag[v] {
			flag[v] = false
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}
