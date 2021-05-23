package main
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0])==0{
		return []int{}
	}
	rows,columns := len(matrix),len(matrix[0])
	//创造一个二维矩阵用来存放这个值是否被访问过
	flag := make([][]bool,rows)
	//要注意给二维切片make空间
	for i := 0;i < rows; i++{
		flag[i]=make([]bool,columns)
	}
	var (
		total = rows * columns
		order = make([]int,total)
		row,column = 0,0
		//逆时针方向
		direction = [][]int{[]int{0,1},[]int{1,0},[]int{0,-1},[]int{-1,0}}
		directionindex = 0
	)
	for i := 0;i < total;i++{
		order[i] = matrix[row][column]
		flag[row][column] = true
		nextRow,nextColumn := row + direction[directionindex][0],column+direction[directionindex][1]
		if nextRow < 0 || nextColumn < 0 || nextColumn >= columns || nextRow >= rows || flag[nextRow][nextColumn]{
			directionindex=(directionindex + 1) % 4
		}
		row += direction[directionindex][0]
		column += direction[directionindex][1]
	}
	return order
}