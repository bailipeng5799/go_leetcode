package main
func removeDuplicates(s string) string {
	length := len(s)
	stack := make([]byte,0) // 因为string类型的字符串中每一个字符是byte类型的
	for i := 0; i < length; i++ {
		//如果栈中有元素并且当前元素等于栈顶元素，那么删除栈顶元素并且忽略此元素
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		}else { //如果没有以上情况那么将其添加到栈中
			stack = append(stack,s[i])
		}
	}
	return string(stack)
}