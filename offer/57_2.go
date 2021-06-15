package main
//双指针+滑动窗口
func findContinuousSequence(target int) [][]int {
	i, j := 1,1
	sum := 0
	ret := make([][]int,0)
	for i <= (target / 2) {
		//如果sum < target 向右移动窗口大小
		if sum < target{
			sum += j
			j++
		}else if sum > target{
			sum -= i
			i++
		}else{
			tmp := make([]int, 0, j-i)
			for m := i; m < j; m++ {
				tmp = append(tmp, m)
			}
			ret = append(ret, tmp)
			sum -= i
			i++
		}
	}
	return ret
}
