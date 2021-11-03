package main
//核心思想就是使用一个map 一次遍历用 map保存之前遍历的
func twoSum( numbers []int ,  target int ) []int {
	// write code here
	hash := map[int]int{}
	for index,value := range numbers {
		if p,ok := hash[target-value];ok{ //在这里查找target-value的值是否在map中存在
			return []int{p+1,index+1}     //如果不存在将其添加到map中
		}
		hash[value] = index
	}
	return []int{}
}
