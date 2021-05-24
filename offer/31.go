package main
func validateStackSequences(pushed []int, popped []int) bool {
	//使用一个新的栈来模仿此程序的入栈和出栈
	//如果判断当前值等于出栈序列第一个值对其进行出栈操作
	//如果不等于继续根据入栈序列进行入栈操作
	//最后判断这个模拟栈的长度是否根据出栈序列全部出栈
	//如果没有说明两个序列不匹配
	stack := []int{}
	for i := 0; i < len(pushed); i++{
		stack = append(stack,pushed[i])
		for len(stack) != 0 && stack[len(stack)-1] == popped[0]{
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(stack) == 0
}
