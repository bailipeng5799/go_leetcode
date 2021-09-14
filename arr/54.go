package main
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0{ //如果矩阵中没有元素那么直接返回
		return []int{}
	}
	res := []int{}
	//top上边界  bottom 下边界   left 左边界   right 右边界
	top, bottom, left, right := 0, len(matrix) - 1, 0, len(matrix[0]) - 1
	//看作是一圈一圈的添加
	for top < bottom && left < right {//循环条件是还存在环，不是一行或者一列
		for i := left; i < right; i++ { res = append(res,matrix[top][i])}
		for i := top; i < bottom; i++ { res = append(res,matrix[i][right])}
		for i := right; i > left; i-- { res = append(res,matrix[bottom][i])}
		for i := bottom; i > top; i-- { res = append(res,matrix[i][left])}
		top++
		bottom--
		left++
		right--
	}
	//top == bottom 说明剩下最后一行
	if top == bottom{
		for i := left; i <= right; i++ { res = append(res,matrix[top][i])}
	}else if left == right { // 说明此时剩下最后一列
		for i := top; i <= bottom; i++ { res = append(res,matrix[i][left])}
	}
	return res

}