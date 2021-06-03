package main

import "fmt"

//构建了一个大顶堆
func HeapSortMax(arr []int,length int){
	if length <= 1{//直接返回进行交换
		return
	}
	depth := length/2 - 1 //非叶子节点的最大下标
	//先构建大顶堆
	for i := depth; i >= 0; i--{
		topmax := i //最大的位置就是i
		leftchild := 2 * i + 1//这个非叶子结点的左右子树下标
		rightchild :=2 * i + 2
		//想要得到这个非叶子节点为根节点的子树的最大值将其放在非叶子节点上
		if arr[leftchild] > arr[topmax] && leftchild <= length - 1{//可能存在没有左子树或者右子树的情况
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
			topmax = rightchild
		}
		if topmax != i {//如果最大值下标被交换了，那么说明将子树最大值的结点与根节点进行交换
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
}
func HeapSort(arr []int) []int{
	length := len(arr)
	for i := 0;i < length; i++{
		lastlen := length - i
		HeapSortMax(arr,lastlen)//重新构建大顶堆
		if i < length {
			//将最大元素沉到数组末端
			arr[0],arr[lastlen-1] = arr[lastlen-1],arr[0]
		}
	}
	return arr
}
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSort(arr))
}



