package main

func minArray(numbers []int) int {
	if len(numbers) == 0{
		return 0
	}
	flag := 0
	for i:=1;i<len(numbers);i++{
		if numbers[i] >= numbers[0]{
			continue
		}
		flag = 1
		return numbers[i]
	}
	if flag == 0{
		return numbers[0]
	}
	return 0

}
func minArray2(numbers []int) int {
	length := len(numbers)
	left,right := 0,length-1
	for left < right{
		mid := left + (right -left) /2
		if numbers[mid] < numbers[right]{
			right=mid
		}else if numbers[mid] > numbers[right]{
			left = mid + 1
		}else{
			right--
		}
	}
	return numbers[left]

}
