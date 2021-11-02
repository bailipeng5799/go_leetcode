package main

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth( a []int ,  n int ,  K int ) int {
	// write code here
	//return heapSort(a,K)
	return quickSortSelectK(a,0,len(a)-1,K)
}

//快排
func quickSortSelectK(arr []int, start,end int, k int)int {
	i, j := start, end
	for i < j {
		for i < j && arr[i] > arr[start] { //要注意满足条件 i < j 并且不能等于因为要得到第k大
			i++
		}
		for i < j && arr[j] <= arr[start] {//这里可以等于是因为只要碰到不大于基准的都可以
			j--
		}
		arr[i], arr[j] = arr[j], arr[i] //将基准进行交换基准前面都是大于基准的值，后面都是小于的
	}
	arr[start],arr[i] = arr[i],arr[start]
	return i  //第i大的值是基准
}
func findk(arr []int,start,end int,k int)int{
	if start <= end {
		tmp := quickSort(arr,start,end,k) //获取到第tmp大的值为tmp
		if tmp == k - 1{ //因为下标从0开始所以当等于k-1的时候返回
			return arr[tmp]
		}else if tmp < k - 1{ //如果tmp比k小那么左节点设置为tmp+1
			return findk(arr,tmp+1,end,k)
		}else { //同理
			return findk(arr,start,tmp-1,k)
		}
	}
	return -1
}
//堆排序
func heapSortMax(arr []int,length int){
	if len(arr) <= 1 {
		return
	}
	depth := length/2
	for i := depth; i >= 0; i--{
		topmax := i
		leftchild := i * 2 + 1
		rightchild := i * 2 + 2
		if leftchild < length && arr[leftchild] > arr[topmax] {
			topmax = leftchild
		}
		if rightchild < length && arr[rightchild] > arr[topmax]{
			topmax = rightchild
		}
		if topmax != i {
			arr[topmax],arr[i] = arr[i],arr[topmax]
		}
	}
}
func heapSort(arr []int,k int) int {
	length := len(arr)
	for i := 0; i < k; i++ {
		lastlen := length - i//因为每一次会得到一个最大值
		heapSortMax(arr,lastlen)
		arr[0],arr[lastlen-1] = arr[lastlen-1],arr[0]
	}
	return arr[length-k]
}


