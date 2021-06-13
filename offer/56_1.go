package main
//本题采用分组异或的方法将两个不同的数字进行分组，在将两个不同组内的所有元素进行异或
func singleNumbers(nums []int) []int {
	k, mask := 0, 1
	//把nums中所有数字进行异或最终得到的结果为两个不同数字的异或结果
	for i := 0; i < len(nums);i++{
		k ^= nums[i]
	}
	//根据异或的性质找到最右一位两个元素不同的位以此作为mask
	//因为连个不同位异或为1，然后再与mask做&操作使mask得到这一位作为区分两个不同元素的标记
	for (mask & k)==0{
		mask <<= 1
	}
	//a，b用来存放两个结果
	a, b := 0, 0
	//根据mask的不同将数字分为两类然后后所有结果进行异或
	for i := 0; i < len(nums); i++{
		if (nums[i] & mask) == 0{
			a ^= nums[i]
		}else{
			b ^= nums[i]
		}
	}
	return []int{a,b}
}