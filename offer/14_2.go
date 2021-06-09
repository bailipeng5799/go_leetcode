package main
func cuttingRope2(n int) int {
	//如果小于3就返回
	if n <= 3{
		return n-1
	}
	res := 1
	//因为n=4的结果就是2*2=4所以直接返回n
	for n>4{
		res = res*3%1000000007
		n -= 3
	}
	//n可能为234
	return res*n%1000000007
}

