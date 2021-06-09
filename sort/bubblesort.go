package main

import "fmt"

func Bubble(arr []int)[]int{
	for i := 0;i < len(arr); i++{
		for j := i + 1;j < len(arr); j++{
			if arr[i] > arr[j]{
				arr[i],arr[j] = arr[j],arr[i]//使i为当前遍历最小值
			}
		}
	}
	return arr
}
//冒泡排序求最大值
func GetMax(arr []int)int{
	for i := 1;i < len(arr); i++{
		if arr[i-1] > arr[i]{
			arr[i-1],arr[i] = arr[i],arr[i-1]
		}
	}
	return arr[len(arr) - 1]
}
func main(){
	arr := []int{1,2,35,6,8,7,88,99,123}
	result := Bubble(arr)
	max := GetMax(arr)
	fmt.Println(result,max)
}
