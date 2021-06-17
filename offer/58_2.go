package main
func reverseLeftWords(s string, n int) string {
	str := []rune(s)
	reserve(str,0,n-1)
	reserve(str,n,len(s) - 1)
	reserve(str,0,len(s) - 1)
	return string(str)
}
func reserve(str []rune,i,j int){
	for ;i < j;i, j = i+1, j - 1{
		str[i],str[j] = str[j],str[i]
	}
}