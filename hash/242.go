package main


func isAnagram(s string, t string) bool {
	//创建一个map来存储字符串中的元素value作为字母出现的个数
	//如果两个字符串长度不相等直接返回
	if len(s) != len(t){
		return false
	}
	flag := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if v, ok := flag[s[i]]; v > 0 && ok{
			flag[s[i]] = v + 1
		}else{
			flag[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++{
		//如果map中的value变为0后在一次需要被删除
		if v, ok := flag[t[i]]; v >= 1 && ok{
			flag[t[i]] = v - 1
		}else{
			return false
		}
	}
	return true
}
