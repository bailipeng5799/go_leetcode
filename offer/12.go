package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m,n := len(board),len(board[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if Search(board,word,i,j,0){
				return true
			}

		}
	}
	return false
}
//k代表单词的第几个字符
func Search(board [][]byte, word string,i,j int,k int)bool{
	if k == len(word){
		return true
	}
	//当i==len(board)时候其实已经超出m的大小因为从0开始
	if i < 0|| j < 0||len(board)==i||len(board[0])==j{
		return false
	}
	//回溯法使用temp临时保存变量
	if board[i][j] == word[k]{
		temp := board[i][j]
		board[i][j] = ' '
		if Search(board,word,i+1,j,k+1) ||Search(board,word,i,j+1,k+1)||
			Search(board,word,i-1,j,k+1) || Search(board,word,i,j-1,k+1){
			return true
		}else{
			board[i][j]=temp
		}
	}
	return false

}

func main(){
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word := "ABCCED"
	result := exist(board,word)
	fmt.Println(result)
}