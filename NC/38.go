package main
//螺旋矩阵
func spiralOrder( matrix [][]int ) []int {
	// write code here
	if len(matrix) == 0 {
		return nil
	}
	res := []int{}
	m,n,mlen,nlen := 0,0,len(matrix)-1,len(matrix[0])-1
	//m,n 代表下一个要遍历的行坐标和纵坐标 mlen,nlen 代表所能遍历到的最后一行或者一列
	for m <= mlen && n <= nlen {
		for i := n; i <= nlen; i++ { //先遍历最上层的一行
			res = append(res,matrix[m][i])
		}
		m++
		for i := m; i <= mlen; i++ { //遍历右边的一行
			res = append(res,matrix[i][nlen])
		}
		nlen--                      //最大列数--
		if m <= mlen {              //如果可以进行行遍历，那么遍历最后一行
			for i := nlen; i >= n; i-- {
				res = append(res,matrix[mlen][i])
			}
		}
		mlen--                      //最大行数--
		if n <= nlen {              //如果可以进行遍历那么遍历最左边一列
			for i := mlen; i >= m; i-- {
				res = append(res,matrix[i][n])
			}
		}
		n++
	}
	return res

}
