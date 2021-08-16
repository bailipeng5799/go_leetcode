package main

import "sort"

/*
	1.先将数组进行排序
	2.双指针法，固定最左元素 k ，双指针分别初始化在k+1 和 切片末尾，所有sum等于0的结果集添加到要返回的切片中
		*当nums[k] > 0 时表明最小的元素已经大于0直接跳出
		*当nums[k]==nums[k-1]时说明和上一次循环相同直接跳过
		*当sum > 0时j--,并且跳过所有重复的nums[j]
		*当sum < 0时i++,并且跳过所有重复的nums[i]
		*当s == 0时，记录组合[k, i, j]至res，执行i += 1和j -= 1并跳过所有重复的nums[i]和nums[j]，防止记录到重复组合。
*/
func threesum(nums []int)[][]int{
	res := [][]int{}
	if len(nums) < 3{
		return res
	}
	sort.Ints(nums)
	for k := 0; k < len(nums)-2; k++{//因为存在三个下标所以需要len(nums) - 2
		if nums[k] > 0{
			break
		}
		if k >0 && nums[k] == nums[k-1]{
			continue
		}
		i,j := k+1, len(nums)-1
		for i < j{
			sum := nums[k] + nums[i] + nums[j]
			if sum > 0{
				//跳过重复的nums[j]
				for i < j && nums[j] == nums[j-1]{
					j--
				}
				j--
			}else if sum < 0{
				//跳过重复的nums[i]
				for i < j && nums[i] == nums[i+1]{
					i++
				}
				i++
			}else{
				tmp := []int{}
				tmp = append(tmp,nums[i],nums[k],nums[j])
				res = append(res,tmp)
				//跳过重复的nums[i]
				for i < j && nums[i] == nums[i+1]{
					i++
				}
				i++
				//跳过重复的nums[j]
				for i < j && nums[j] == nums[j-1]{
					j--
				}
				j--
			}

		}
	}
	return res
}
