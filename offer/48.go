package main
func lengthOfLongestSubstring(s string) int {
	hashTable:=map[byte]int{}
	// i赋一个不可能的初值，为了更新
	i,res:=-1,0
	for j:=0;j<len(s);j++{
		//判断是否第二次出现
		if p,ok:=hashTable[s[j]];ok{
			// i存的是这个字符上次出现的位置
			i=max(i,p)
		}
		//新增或更新字符的位置
		hashTable[s[j]]=j
		// j-i为字符位置和上一次位置的差 ,每次都比较这个差值
		res=max(res,j-i)
	}
	return res
}

func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}
