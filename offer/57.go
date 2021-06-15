package main
func twoSum(nums []int, target int) []int {
	//左右两个指针判断两个指针之和与target的大小
	//如果sum大于target那么将右指针左移
	//如果sum小于target那么将左指针右移
	p, q := 0, len(nums) - 1
	for p < q {
		s := nums[q] + nums[p]
		if s < target{
			p++
		}else if s > target{
			q--
		}else{
			return []int{nums[q],nums[p]}
		}
	}
	return []int{}
}