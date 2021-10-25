package main

func HeapSortMin(input []int,length int)[]int{
	if length <= 1{
		return input
	}
	depth := length / 2 -1
	for i := depth;i >= 0;i--{
		topmin := i // 最小的位置就是i
		leftchild := 2 * i + 1//这个非叶子结点的左右子树下标
		rightchild :=2 * i + 2
		if input[leftchild] < input[topmin] && leftchild <= length -1{
			topmin = leftchild
		}
		if input[rightchild] < input[topmin] && rightchild <= length -1{
			topmin = rightchild
		}
		if topmin != i{
			input[topmin],input[i]=input[i],input[topmin]
		}
	}
	return input
}
func GetLeastNumbers_Solution( input []int ,  k int ) []int {
	for i := 0;i < k;i++{
		lastlen := len(input) - i
		HeapSortMin(input,lastlen)
		if i < len(input) {
			//将最小元素沉到数组末端
			input[0],input[lastlen-1] = input[lastlen-1],input[0]
		}
	}
	return input[len(input)-k:]
}