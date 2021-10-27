package main

var stack1 [] int
var stack2 [] int
/*
   push操作就直接往stack1中push， pop操作需要分类一下：如果stack2为空，那么需要将stack1中的数据转移到stack2中，
   然后在对stack2进行pop，如果stack2不为空，直接pop就ok。
*/
func Push(node int) {
	stack1 = append(stack1,node)
}

func Pop() int{
	if len(stack2) != 0 {
		res := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		return res
	}

	for i := len(stack1)-1; i > 0; i-- {
		stack2 = append(stack2,stack1[i])
	}
	res := stack1[0]
	stack1 = []int{}

	return res
}
