package main

func findKthLargest(nums []int, k int) int {
	res := HeapSort(nums,k)
	return res[len(nums)-k]
}
func HeapSort(arr []int, k int) []int {
	length := len(arr)
	for i := 0; i < k; i++ {
		lastlen := length - i
		HeapSortMax(arr,lastlen)
		arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
	}
	return arr
}
func HeapSortMax(arr []int,length int){
	if length <= 1 {
		return
	}
	tmp := length/2 - 1 //最后一个非叶子节点
	for i := tmp; i >= 0; i-- {
		topmax := i
		leftChild := i * 2 + 1
		rightChild := i * 2 + 2
		if leftChild < length && arr[leftChild] > arr[topmax] {
			topmax = leftChild
		}
		if rightChild < length && arr[rightChild] > arr[topmax] {
			topmax = rightChild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
}



func findKthLargest(nums []int, k int) int {
	return Partition(nums,k)
}

func Partition(nums []int,k int)int{    //快排
	i,j:=0,len(nums)-1
	pivot:=nums[0]
	for i<j{
		for i<j && nums[j]<pivot{
			j--
		}
		if i<j{
			nums[i],nums[j]=nums[j],nums[i]
			i++
		}
		for i<j && nums[i]>pivot{
			i++
		}
		if i<j{
			nums[i],nums[j]=nums[j],nums[i]
			j--
		}
	}

	if i+1==k{
		return pivot  //   位子恰好就是K-1返回结果
	}else if i+1<k{
		return Partition(nums[i+1:],k-i-1)       //以pivot为断点，在后半部分数组中进行查找
	}else{
		return Partition(nums[:i],k)    //在前半部分数组中进行查找
	}
}
