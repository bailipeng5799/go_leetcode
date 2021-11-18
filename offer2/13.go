package main
func movingCount(m int, n int, k int) int {
	board := make([][]bool,m)
	for i := 0; i < m; i++ {
		board[i] = make([]bool,n)
	}
	res := 0
	var Search func(board [][]bool,i,j int)bool
	Search = func(board [][]bool,i,j int)bool{
		if i >= m || j >= n || i < 0 || j < 0 || board[i][j] || company(i,j) > k {
			return false
		}
		board[i][j] = true
		res++
		if Search(board,i+1,j) || Search(board,i-1,j) || Search(board,i,j+1) || Search(board,i,j-1) {
			return true
		}

		return false

	}
	if Search(board,0,0) {
		return res
	}
	return res
}
func company(i,j int)int {
	tmp := 0
	for i > 0 {
		tmp += i % 10
		i /= 10
	}
	for j > 0 {
		tmp += j % 10
		j /= 10
	}
	return tmp
}