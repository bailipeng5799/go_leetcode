package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	temp:=math.Pow(10.0,float64(n))
	result := []int{}
	for i := 1;i < int(temp); i++{
		result = append(result,i)
	}
	return result
}
func main(){
	result := printNumbers(1)
	fmt.Println(result)
}