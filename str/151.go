package main
func reverseWords(s string) string {
	//主要思想是利用两个指针来消除空格
	//fastIndex指针碰到空格就直接++，slowIndex作为每一次替换的对象
	slowIndex, fastIndex := 0, 0
	str := []byte(s)
	var reverse func(start,end int)
	reverse = func(start, end int){
		for start < end {
			str[start], str[end] = str[end], str[start]
			start += 1
			end -= 1
		}
	}
	//删除头部空格主要思想是碰到空格fastIndex++
	for len(str) > 0 && fastIndex < len(str) && str[fastIndex] == ' '{
		fastIndex++
	}
	//删除单词中间空格主要思想是覆盖的思想 使用slowIndex和fastIndex指针来覆盖多余的空格 fastIndex碰到两个连续空格就cotinue
	for ; fastIndex < len(str); fastIndex++{
		if fastIndex - 1 > 0 && str[fastIndex] == str[fastIndex - 1] && str[fastIndex] == ' '{
			continue
		}
		str[slowIndex] = str[fastIndex]
		slowIndex++
	}
	//删除末尾的多余空格
	if slowIndex - 1 > 0 && str[slowIndex - 1] == ' '{
		str = str[:slowIndex-1]
	}else{
		str = str[:slowIndex]
	}
	//先翻转全部字符串 然后翻转单个单词
	reverse(0,len(str)-1)
	//翻转单个单词 i表示单词开始位置，j表示单词结束位置
	i := 0
	for i < len(str) {
		j := i
		for ;j<len(str) && str[j] != ' ';j++ {

		}
		reverse(i,j-1)
		i=j//此时i==' '
		i++
	}
	return string(str)

}
