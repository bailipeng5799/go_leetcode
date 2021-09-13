package main
func search(nums []int, target int) int {
	/*
	   for left <= right 要使用 <= ，因为left == right是有意义的，所以使用 <=
	   if nums[middle] > target, right 要赋值为 middle - 1，因为当前这个nums[middle]
	   一定不是target，那么接下来要查找的左区间结束下标位置就是 middle - 1
	*/
	left, right := 0, len(nums) - 1 //定义target在左闭右闭的区间里，[left, right]
	for left <= right {				//当left==right，区间[left, right]依然有效，所以用 <=
		mid := left + (right-left) / 2 //防止溢出 等同于(left + right)/2
		if nums[mid] == target{
			return mid
		}else if(nums[mid] > target){
			right = mid - 1  // target 在左区间，所以[left, middle - 1]
		}else{
			left = mid +1 	// target 在右区间，所以[middle + 1, right]
		}
	}
	return -1
}
func search2(nums []int, target int) int {
	left, right := 0, len(nums) //定义target在左闭右开的区间里，即：[left, right)注意两个right的取值不同
	for left < right {			//当left == right时 区间[left,right)是无效的空间  所以使用<
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right = mid //target在左区间在 [left,mid)中
		}else {
			left = mid + 1 //target在右区间[mid+1,right)
		}
	}
	return -1
}