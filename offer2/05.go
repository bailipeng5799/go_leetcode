package main
func replaceSpace(s string) string {
	//for range
	var res []int32
	for _,value := range s { //这里从s中获取的value为rune类型也就是int32
		if value == ' ' {
			res = append(res,'%','2','0') //''表示rune类型，""表示string类型
		}else {
			res = append(res,value)
		}
	}
	return string(res)
}