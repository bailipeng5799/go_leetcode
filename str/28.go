package main
//KMP算法重要
func strStr(haystack string, needle string) int {
	//两次循环暴力解法
	left, right := 0, 0
	countIndex := 0
	for left <= len(haystack) - len(needle) {
		for countIndex < len(needle) && haystack[right] == needle[countIndex] {
			right++
			countIndex++
		}
		if countIndex == len(needle) {
			return left
		}
		left++
		right = left
		countIndex = 0
	}
	return -1
}
func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	j := 0
	next := make([]int,len(needle))
	getNext(next,needle)//获取next数组，存储每一个index最长前后缀相等的长度
	for i := 0; i < len(haystack); i++{
		for j > 0 && haystack[i] != needle[j]{//当前位置字符不相等从next中得到最长相当等的前缀字符串的index
			//如果还是不相等就一直向前返回最多可以返回到needle的头部
			j = next[j-1]
		}
		if haystack[i] == needle[j]{ //如果相等j++最长前后缀相等的长度加1
			j++
		}
		if j == len(needle){
			return i - len(needle) + 1 //因为当前i为比较后的下标，结果需要返回第一个位置
		}
	}
	return -1
}
func getNext(next []int, s string) { //获取前缀表
	//i为后缀的末尾字符  j为前缀的末尾字符  j也等于当前下标的最长前后缀字符串长度
	j := 0
	next[0] = j     //初始化next
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j]{
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}