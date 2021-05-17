package main

import "fmt"
//递归
func myPow(x float64, n int) float64 {
	if x == 0{
		return 0
	}
	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}
	if n < 0{
		n *= -1
		x=1/x
	}
	temp:=myPow(x,n/2)
	if n % 2 == 0{
		return temp*temp
	}
	return x*temp*temp
}
func main() {
	result:= myPow(2.00000,10)
	fmt.Printf("%v",result)
}
