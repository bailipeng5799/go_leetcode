package main

func getLeastNumbers(arr []int, k int) []int {
	res := QuickSort(arr)
	return res[:k]
	//res := HeapSort(arr,k)
	return res

}
//快排topk
func QuickSort(nums []int) []int{
	if len(nums) <= 1{
		return nums
	}
	splitdata := nums[0]
	low := make([]int, 0)
	mid := make([]int, 0)
	hight := make([]int, 0)
	mid = append(mid,splitdata)
	for i := 1; i < len(nums); i++{
		if nums[i] > splitdata{
			hight = append(hight, nums[i])
		} else if nums[i] < splitdata{
			low = append(low, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	return append(append(low, mid...), hight...)
}

func HeapSortMin(arr []int, length int){
	// length := len(arr)
	if length <= 1 {
		return
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topmin := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] < arr[topmin] { //防止越过界限
			topmin = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] < arr[topmin] { //防止越过界限
			topmin = rightchild
		}
		if topmin != i {
			arr[i], arr[topmin] = arr[topmin], arr[i]
		}
	}
}
func HeapSort(arr []int,k int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMin(arr, lastlen)
		if i < length && i <= k{
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr[length-k:]
}


