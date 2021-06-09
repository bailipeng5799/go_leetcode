package main
//摩尔投票法
//按需遍历如果有两个不一样的就将投票数抵消，最后票数不为0的x为所求
func majorityElement(nums []int) int {
	x  := nums[0]
	votes := 1
	for i := 1;i<len(nums);i++{
		if x == nums[i]{
			votes++
		}else{
			votes--
		}
		if votes == 0{
			x=nums[i+1]
		}
	}
	return x
}
