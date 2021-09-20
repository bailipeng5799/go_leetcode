package main

func reverseLeftWords(s string, n int) string {
	tmp := []byte(s)
	xbxreverse(tmp,0,n-1)
	xbxreverse(tmp,n,len(tmp)-1)
	xbxreverse(tmp,0,len(tmp)-1)
	return string(tmp)
}
func xbxreverse(tmp []byte,left, right int){
	for left < right {
		tmp[left],tmp[right] = tmp[right],tmp[left]
		left++
		right--
	}
}
