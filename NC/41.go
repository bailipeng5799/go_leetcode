package main

func maxLength( arr []int ) int {
	//write code here
	left, right, tmp, res := 0, 0, 0, 0

	for left < len(arr) {
		hash := map[int]bool{}
		right = left
		for right < len(arr) {
			if _,ok := hash[arr[right]]; ok{

				break
			}
			hash[arr[right]] = true
			tmp++
			right++
		}

		left++
		res = max(res,tmp)
		tmp = 0

	}
	return res
}
func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}
