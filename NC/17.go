package main
func getLongestPalindrome( A string ,  n int ) int {
	// write code here
	maxlen := 0
	i, j := 0, 0
	for i < n && j < n {
		itemp := i //暂时存储ij
		jtemp := j
		templen := 0
		for itemp >= 0 && jtemp < n && A[itemp] == A[jtemp] { //为什么要大于0是因为肯呢个遍历到第一个元素之后
			//两端元素依然相等，这是就会不满足条件
			itemp--
			jtemp++
		}
		templen = jtemp - itemp - 1 //计算回文子串长度
		if templen > maxlen {
			maxlen = templen
		}
		if i == j { //最长回文子串长度为偶数 同时从中间两个元素遍历
			j++
		}else { //最长回文子串长度为奇数，这时候从中间相同一个元素向两边遍历
			i++
		}
	}
	return maxlen
}
