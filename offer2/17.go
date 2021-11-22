package main

import "strconv"

func printNumbers(n int) []int {
	res := []int{}
	var dfs func(index int,num []byte,digit int)
	dfs = func (index int,num []byte,digit int) {
		if index == digit { //说明这个数字的最后一位已经被存储num中
			tmp, _ := strconv.Atoi(string(num))
			res = append(res,tmp)
			return
		}
		for i := '0'; i <= '9'; i++ {
			num[index] = byte(i)
			dfs(index+1,num,digit)
		}
	}
	for digit := 1; digit <= n; digit++ { //代表位数
		for first := '1'; first <= '9'; first++ { //第一位从1—9开始
			num := make([]byte,digit)
			num[0] = byte(first)                  //创建一个num切片用来暂时存放数字
			dfs(1,num,digit)
		}
	}
	return res

}


