package main
func reverseStr(s string, k int) string {
	tmp := []byte(s)
	for i := 0;i < len(s);i += 2*k {
		if i + k <= len(s){
			//说明剩余字符小于2k大于等于k个
			//这时反转前k个字符
			reverse(tmp[i:i+k])
		}else{
			//说明剩余字符小于k个
			//反转全部的字符
			reverse(tmp[i:len(s)])
		}
	}
	return string(tmp)
}
func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left],b[right] = b[right],b[left]
		left++
		right--
	}
}
