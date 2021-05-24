package main
//使用一个辅助栈用来存放每一次入栈时当前栈内的最小元素
type MinStack struct {
	stack []int
	min_stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min_stack:[]int{1<<63-1},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack,x)
	this.min_stack = append(this.min_stack,min(x,this.min_stack[len(this.min_stack)-1]))
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.min_stack = this.min_stack[:len(this.min_stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) Min() int {
	return this.min_stack[len(this.min_stack)-1]
}
func min(a,b int)int{
	if a > b{
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
