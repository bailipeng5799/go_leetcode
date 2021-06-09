package main

import "fmt"

func fib(n int) int {
	nums:=make([]int,n+1)
	if n == 0{
		return 0
	}
	if n == 1{
		return 1
	}
	nums[0],nums[1]=0,1
	for i:=2;i <= n; i++{
		nums[i]=(nums[i-1]+nums[i-2])% 1000000007
	}
	return nums[n]
}
func fib2(n int) int {
	f0,f1 := 0, 1
	for i := 0;i < n;i++{
		f0,f1=f1,(f0+f1)%1000000007
	}
	return f0
}
func main(){
	result := fib2(95)
	fmt.Println(result)
}