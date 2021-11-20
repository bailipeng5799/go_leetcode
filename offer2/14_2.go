package main
func cuttingRope2(n int) int {
	if n <= 3 {
		return n-1
	}
	res := 1
	//n = 4时最大值就是 2 * 2 = 4 所以利用4作为条件
	for n > 4 {
		res = res * 3 %(1e9+7)
		n -= 3
	}
	//n 可能为 1 2 3 4的一种
	return res * n % (1e9+7)
}

