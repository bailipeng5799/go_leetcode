package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0])==0{

		return false
	}
	if target < matrix[0][0] ||target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	i,j:=0,len(matrix[0])-1
	for i < len(matrix) && j >= 0{
		if matrix[i][j] == target{
			return true
		}
		if matrix[i][j] < target{
			i++
			continue
		}
		j--
	}
	fmt.Println(1)
	return false
}
func main(){
	//matrix:=[][]int{{1, 4, 7, 11, 15},{2, 5, 8, 12, 19},{3, 6, 9, 16, 22},{10, 13, 14, 17, 24},{18, 21, 23, 26, 30}}
	matrix2:=[][]int{{-5}}
	result := findNumberIn2DArray(matrix2,-5)
	fmt.Println(result)
}