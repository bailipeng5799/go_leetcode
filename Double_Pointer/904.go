package main
//滑动窗口+map
func totalFruit(fruits []int) int {
	count := 0
	//左指针，右指针
	left, right := 0, 0
	species := make(map[int]int)//策划构建一个map用来存储水果的种类
	for right < len(fruits) {//右指针还在界限之类
		species[fruits[right]]++ //每次将右边界所指向的水果种类++
		if len(species) > 2{ //如果map中元素个数大于2
			species[fruits[left]]--//先将窗口中左边界的种类减少1
			if species[fruits[left]] == 0{ //判断个数如果没有此类元素了就从map中删除
				delete(species,fruits[left])
			}
			left++//左边界右移
		}
		tmpCount := right - left + 1//查看窗口大小也就是篮子中水果的数量
		if tmpCount > count { //更新结果
			count = tmpCount
		}
		right++
	}
	return count
}