package main
func repeatedSubstringPattern(s string) bool {
	//KMP
	//数组长度减去最长相同前后缀的长度相当于是第一个周期的长度，也就是一个周期的长度，
	//如果这个周期可以被整除，就说明整个数组就是这个周期的循环。
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int,n)
	next[0] = j
	for i := 1; i < n; i++{
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j]{
			j++
		}
		next[i] = j
	}
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}