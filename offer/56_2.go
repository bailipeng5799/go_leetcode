package main
//map确定次数
func singleNumber(nums []int) int {
	count := make(map[int]int,len(nums)/3 + 1)
	for _ ,value := range nums{
		count[value]++
	}
	for index ,value := range count{
		if value == 1{
			return index
		}
	}
	return 0
}
