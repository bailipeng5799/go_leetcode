package main

func generateMatrix(n int) [][]int {
	top, bottom, left, right := 0, n-1, 0, n-1 //上下左右边界
	num := 1
	//申请空间
	res := make([][]int,n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	//和螺旋数组1略微有一点不同之处在于对于中间元素的处理
	//本题是直接通过循环进行处理，每一次遍历完后直接对边界进行操作
	for num <= n * n{
		for i := left; i <= right; i++{
			res[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++{
			res[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i--{
			res[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i--{
			res[i][left] = num
			num++
		}

		left++
	}
	return res
}

