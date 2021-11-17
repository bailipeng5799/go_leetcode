package main
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	f0, f1 := 1, 2
	for i := 3; i < n; i++ {
		f0, f1 = f1,(f1+f0)%(1e9+7)
	}
	return f1
}
