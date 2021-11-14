package main
func findRepeatNumber(nums []int) int {
	hash := make(map[int]bool,len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return nums[i]
		}
		hash[nums[i]]=true
	}
	return -1
}
func findRepeatNumber2(nums []int) int {
	i := 0
	for i < len(nums){ //一直进行交换直到索引下标为0时值同样为0
		if nums[i] == i{ //如果索引值和所对应的值相等直接跳过
			i++
			continue
		}
		if nums[i] == nums[nums[i]] { //如果查询到对应位置上的值和当前值相等则元素重复
			return nums[i]
		}
		nums[i],nums[nums[i]] = nums[nums[i]],nums[i] //和以nums[i]为索引的地方的值进行交换
	}
	return -1
}