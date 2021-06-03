package main

import (
	"fmt"
)

func countDigitOne(n int) int{
	if n == 0 {
		return 0
	}
	if n < 10 {
		return 1
	}
	cur := 1//cur = 10^i
	i:= 0
	res := 0
	//i得到他是几位数
	for cur < n{
		cur*=10
		i++
	}
	cur /= 10//得到的cur为小于这个数的最小十次方
	//如果他是10的倍数
	if cur * 10 == n{
		return i * cur + 1
	}
	i--
	for cur > 0 {
		t := n / cur
		n = n % cur
		if t == 1{
			res += i * cur / 10 + n + 1
		}else{
			res += t * i*cur/10 + cur
		}
		for cur > n{
			cur /= 10
			i--
		}
	}
	return res
}
func main(){
	res := countDigitOne(12)
	fmt.Println(res)
}
