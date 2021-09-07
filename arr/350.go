package main

import "sort"

//350. 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1{
		//遍历nums1，将数据添加到map中
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2{
		//遍历nums2，将map中存在的数据进行-1
		//这里直接在nums2中进行赋值不需要重新申请空间，使用k这个中间变量来记录交集的个数
		//使用map是因为可能存在多个相同的元素
		if m0[v] > 0{
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}
//如果给定的数组已经排好序呢
//设定两个为0的指针，比较两个指针的元素是否相等。 如果指针的元素相等，
//我们将两个指针一起向后移动，并且将相等的元素放入空白数组。
//我们的指针分别指向第一个元素，判断元素相等之后，将相同元素放到空白的数组。
//如果两个指针的元素不相等，我们将小的一个指针后移。
//我们指针移到下一个元素，判断不相等之后，将元素小的指针向后移动，继续进行判断。
//直到任意一个数组终止。
func intersect2(nums1 []int,nums2 []int) []int{
	i,j,k := 0,0,0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2){
		if nums1[i] > nums2[j]{
			j++
		}else if nums1[i] < nums2[j]{
			i++
		}else{
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}