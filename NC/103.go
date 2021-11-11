package main
// 反转字符串
func solve1( str string ) string {
	// write code here
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	for i := 0; i < len(str)/2; i++ {
		tmp[i],tmp[len(str)-i-1] = tmp[len(str)-i-1],tmp[i]
	}
	return string(tmp)
}