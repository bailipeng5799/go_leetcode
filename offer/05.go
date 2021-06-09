package main

import "fmt"
//双引号操作byte类型  uint8
//单引号操作rune类型  int32
func replaceSpace(s string) string {
	str := make([]rune,len(s)*3)
	i := 0
	for _,value := range s{
		if value != ' '{
			str[i]=value
			i++
		}else{
			str[i]='%'
			str[i+1]='2'
			str[i+2]='0'
			i+=3
		}
	}
	return string(str)[:i]
}
func replaceSpace2(s string) string{
	var result []rune
	for _,value := range s{
		if value ==32{
			result=append(result,'%','2','0')
		}else{
			result=append(result,value)
		}
	}
	return string(result)
}
func main(){
	s := "We are happy."
	result := replaceSpace(s)
	fmt.Println(result)
}