package main
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0{
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n *= -1
		x = 1/x
	}
	res := 1.0
	for n >= 1{
		if n & 1 == 1 {
			res *= x
			n--
		} else {
			x *= x
			n = n >> 1 //   ➗2
		}
	}
	return res
}
