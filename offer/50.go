package main
func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	count := make(map[byte]int,0)
	for i := 0;i < len(s); i++{
		count[s[i]]++
	}
	for i := 0;i < len(s); i++{
		if count[s[i]] == 1{
			return s[i]
		}
	}
	return ' '
}
