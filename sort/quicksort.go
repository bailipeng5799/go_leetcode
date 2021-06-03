package main

import "fmt"

/*
1.从数列中挑出一个元素，称为 "基准"（pivot）;
2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/
func QuickSort(nums []int) []int{
	if len(nums) <= 1{
		return nums
	}
	splitdata := nums[0]
	low := make([]int, 0, 0)//存放比它小的
	mid := make([]int, 0, 0)//相等的
	hight := make([]int, 0, 0)//大的
	mid = append(mid,splitdata)
	for i := 1; i < len(nums); i++{
		if nums[i] < splitdata{
			low = append(low, nums[i])
		}else if nums[i] > splitdata{
			hight = append(hight, nums[i])
		}else {
			mid = append(mid, nums[i])
		}
	}
	low, hight = QuickSort(low),QuickSort(hight)//递归调用
	return append(append(low, mid...), hight...)//返回拼接好的切片
}
//第二中
func QuickSort2(nums []int, left,right int) []int{
	if left < right {
		partitionIndex := partition(nums, left, right)
		QuickSort2(nums, left, partitionIndex-1)
		QuickSort2(nums, partitionIndex+1, right)
	}
	return nums
}
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1//最终用来交换基准和分割切片
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
	fmt.Println(QuickSort2(arr,0,len(arr)-1))
}