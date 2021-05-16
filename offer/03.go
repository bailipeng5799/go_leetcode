package main

import "fmt"

func findRepeatNumber(nums []int) int {
	temp := [100000]int{}
	for _ ,value := range nums{
		temp[value]++
		if temp[value]==2{
			return value
		}
	}
	return -1
}
//原地置换
func findRepeatNumber2(nums []int) int {
	for index,value := range nums{
		if index == value {
			continue
		}
		if nums[value]==value{
			return value
		}
		nums[index],nums[value]=nums[value],nums[index]
	}
	return -1
}

func main(){
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	result := findRepeatNumber2(nums)
	fmt.Println(result)
}
