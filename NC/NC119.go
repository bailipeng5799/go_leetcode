package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */

func HeapSortMin(input []int,length int){
	if len(input) <= 1 {
		return
	}
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		topmin := i
		leftchild := i*2 + 1
		rightchild := i*2 + 2
		if leftchild <= length-1 && input[leftchild] < input[topmin] {
			topmin = leftchild
		}
		if rightchild <= length-1 && input[rightchild] < input[topmin] {
			topmin = rightchild
		}
		if topmin != i {
			input[topmin],input[i] = input[i],input[topmin]
		}
	}
}
func GetLeastNumbers_Solution( input []int ,  k int ) []int {
	for i := 0; i < k; i++ {
		lastlen := len(input) - i
		HeapSortMin(input,lastlen)//将最小的元素放到了 0 这个位置上
		if i < len(input) {
			input[0],input[lastlen-1] = input[lastlen-1],input[0]
		}
	}
	return input[len(input)-k:]
}