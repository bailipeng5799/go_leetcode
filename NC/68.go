package main
//跳台阶
func jumpFloor( number int ) int {
	if number <= 2 {
		return number
	}
	res := make([]int,number+1) //因为0元素已经被占用
	res[0],res[1],res[2] = 0,1,2
	for i := 3; i <= number; i++{
		res[i] = res[i-1]+res[i-2]
	}
	return res[number]
	//     a,b := 1,1
	//     for i := 1;i < number; i++{
	//         a,b = b,a+b
	//     }
	//     return b
}
