package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	count := len(strs)//一共有多少个字符串
	tmp := strs[0]//使用tmp来存储默认最长公共前缀默认为strs[0]
	for i := 1;i < count;i++{
		tmp = lcp(strs[i],tmp)//返回两个参数的公共前缀
		if len(tmp) == 0{
			break
		}
	}
	return tmp
}
//构建一个函数来比较两个字符串的最长公共前缀并且返回
func lcp(str1 ,str2 string)string{
	count := min(len(str1),len(str2))//记录最小的字符串的长度
	index := 0//临时下标，用来返回公共前缀
	for index < count && str1[index] == str2[index]{
		index++
	}
	return str1[:index]
}
func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}