package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n, mlen, nlen := 0, 0, len(matrix)-1, len(matrix[0])-1
	//m,n代表当前行和当前列 mlen，nlen代表最大行和最大列
	res := []int{}
	for m <= mlen && n <= nlen {
		for i := n; i <= nlen; i++ {
			res = append(res,matrix[m][i]) //遍历第m行的范围内的每一列
		}
		//结束之后第 m 行遍历结束，
		m++
		for i := m; i <= mlen; i++ {
			res = append(res,matrix[i][nlen]) //遍历最大列的范围内的每一行元素
		}
		//结束之后最大列 nlen遍历结束
		nlen--
		if m <= mlen {
			for i := nlen; i >= n; i-- {
				res = append(res,matrix[mlen][i]) //遍历最大行范围内的每一个元素
			}
		}
		//结束之后最大行 mlen 遍历结束
		mlen--
		if n <= nlen {
			for i := mlen; i >= m; i-- {
				res = append(res,matrix[i][n])
			}
		}
		//结束之后第n列遍历结束
		n++
	}
	return res
}