package main
func exist(board [][]byte, word string) bool {
	if len (word) > len(board)*len(board[0]) {
		return false
	}
	n, m := len(board), len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if Search(board,word,i,j,0) {
				return true
			}
		}
	}
	return false
}
func Search(board [][]byte, word string,i, j int, k int) bool {
	if k == len(word) {
		//k代表单词的第几个字符
		return true
	}
	if i < 0 || j < 0 || len(board) == i || len(board[0]) == j {
		return false
	}
	if board[i][j] == word[k] {
		tmp := board[i][j]
		board[i][j] = ' '//表示为空代表已经使用
		if Search(board,word,i+1,j,k+1)|| Search(board,word,i,j+1,k+1)||Search(board,word,i-1,j,k+1)||Search(board,word,i,j-1,k+1) {
			return true
		}else {
			board[i][j] = tmp
		}

	}
	return false
}