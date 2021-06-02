package main
func findNthDigit(n int) int {
	tmp,tmpNum,count := 1,1,9//tmp表示位数，tmpNum表示10^x count 表示一位数有九位
	for n > count {
		n -= count
		tmp++
		tmpNum *= 10
		count = 9 * tmp * tmpNum//表示某一数有多少位
	}
	num := tmpNum + (n - 1)/tmp    //n-1是根据找规律得到的
	index := (n - 1) % tmp		//对这个数是几位数进行取模
	numstr := strconv.Itoa(num)
	return int(numstr[index]-'0')
}
