package main
func replaceSpace(s string) string {
	//for range  中拿出的元素是 rune类型的元素

	var res []rune
	for _,value := range s {
		if value == ' '{
			res = append(res,'%','2','0')
		}else{
			res = append(res,value)
		}
	}
	return string(res)
}
