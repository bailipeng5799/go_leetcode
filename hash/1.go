package main
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	//一次遍历每一次将元素值和坐标存入map中
	for index,value := range nums{
		if p,ok := hash[target-value];ok{
			return []int{index,p} //如果map中存在target-value 的值那么直接返回所对应的坐标
		}
		hash[value]=index

	}
	return []int{}
}
