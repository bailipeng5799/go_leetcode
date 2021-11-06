package main
/*
 括号序列
	给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
	括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

	数据范围：字符串长度 0\le n \le 100000≤n≤10000
	要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
	输入："["
	返回值：false
*/
func isValid( s string ) bool {
	// write code here
	if len(s) % 2 == 1 {
		return false
	}
	var stack []int32
	for _,c := range s{
		switch c {
		case '(','[','{':
			stack = append(stack,c)
		case ']':
			if len(stack) == 0 {
				return false
			}else if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		case ')':
			if len(stack) == 0 {
				return false
			}else if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		case '}':
			if len(stack) == 0 {
				return false
			}else if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		default :
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
