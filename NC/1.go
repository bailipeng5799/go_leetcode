package main
//大数加法
func solve( s string ,  t string ) string {
	// 先将长的数字放在 t 中 短的数字放在 m中
	//主题思想就是利用加法运算先从最小位开始加起
	m := len(s)
	n := len(t)
	if m > n {
		t, s = s, t
		m, n = n, m
	}
	ret := make([]byte, n+1)
	carry := 0 //表示需要jin的位比如7+8 jin 1
	for i := 0; i < n; i++ {
		longN := int(t[n-i-1] - '0')
		shortN := 0 //默认值为0如果没有的化这一位为0
		if m-i-1 >= 0 {
			shortN = int(s[m-i-1] - '0')
		}
		result := longN + shortN +carry
		carry = result / 10
		result = result % 10
		ret[n-i] = byte(result + '0')
	}
	//最后判断carry是否为1
	if carry == 1{
		ret[0] = '1'
		return string(ret)
	}
	return string(ret[1:])
}
