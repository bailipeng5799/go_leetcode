package main
func findNumberIn2DArray(matrix [][]int, target int) bool {
	//从左下角开始遍历
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j:= len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if target == matrix[i][j] {
			return true
		}else if target > matrix[i][j] { //如果要查找的数大于当前坐标的值i++
			j++
		}else if target < matrix[i][j] { //如果要查找的数小于当前坐标的值j--
			i--
		}

	}
	return false
}
