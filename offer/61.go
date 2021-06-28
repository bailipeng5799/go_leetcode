package main

import "sort"

func isStraight(nums []int) bool {
	//利用五个连续数字的茶之必定小于5得到答案
	//nums 转化为递增序列
	sort.Ints(nums)
	//记录每个数字之间差值的和
	sub:=0
	for i:=0;i<4;i++{
		if nums[i]==0{
			continue
		}
		//如果扑克牌（非0数字）重复，不是顺子
		if nums[i]==nums[i+1]{
			return false
		}
		//差值记录
		sub+=nums[i+1]-nums[i]
	}
	//总的差值小于5（或sub<=4）则为顺子
	return sub<5
}
