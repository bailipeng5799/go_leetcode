package main

import "container/heap"

//构建小顶堆
type IHeap [][2]int
//因为自己初始化了一个小顶堆，所以需要实现相关接口
func (h IHeap) Len()int {
	return len(h)
}
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }//用来决定升序还是降序
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) { //入堆
	*h = append(*h,x.([2]int))
}
func (h *IHeap) Pop() interface{} { //删除堆顶元素并且返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	//记录每个元素出现的次数
	for _ ,item := range nums {
		hash[item]++
	}
	h := &IHeap{} //初始化一个小顶堆
	heap.Init(h)
	for key, value := range hash {
		heap.Push(h,[2]int{key,value})//先push元素，如果发现当前长度大于k删除最小的元素也就是堆顶的元素
		if h.Len() > k{
			heap.Pop(h)
		}
	}
	res := make([]int,k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}//拿到小顶堆中最大的两个元素，也就是距离堆顶最远的两个元素
	return res
}