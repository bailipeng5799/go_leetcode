package main
func canConstruct(ransomNote string, magazine string) bool {
	//使用一个map进行简单的操作
	//先将magazine的元素存入到map中
	//第二次遍历ransomNote将map中的元素每次--
	//如果遍历中不存在此元素或者当前元素个数=0 直接返回false
	record := make(map[rune]int)
	for _, value := range magazine {
		if v, ok := record[value];ok{
			record[value] = v + 1
		}else{
			record[value] = 1
		}
	}
	for _, ransomValue := range ransomNote {
		if v, ok := record[ransomValue];ok && v > 0 {
			record[ransomValue] = v - 1
		}else{
			return false
		}
	}
	return true
}
func canConstruct2(ransomNote string, magazine string) bool {
	//因为只存在26种字符
	//所以使用int类型的切片既可以解决
	record := make([]int,26)
	for _, v := range magazine {
		//rune - 'a' 作为下标代表这个元素的数量
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0{
			return false
		}
	}
	return true
}