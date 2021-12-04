package main
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	hash := map[byte]int{}
	maxvalue := 0
	for left < len(s) {
		if left != 0 {
			delete(hash,s[left-1])//删除上一个元素
		}
		for right < len(s) && hash[s[right]] == 0 {
			hash[s[right]] = 1
			right++
		}
		maxvalue = max(maxvalue,right-left)
		left++
	}
	return maxvalue
}
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
