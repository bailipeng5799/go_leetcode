package main
//如果x知道
func fib(n int) int {
	if n == 0{
		return 0
	}
	if n < 2 {
		return 1
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % (1e9 + 7)
	}
	return r


}
