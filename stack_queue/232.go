package main
type MyQueue struct {
	stack []int//输入栈
	back  []int//输出栈
}


func Constructor() MyQueue {
	return MyQueue{
		stack : make([]int,0),
		back : make([]int,0),
	}
}


func (this *MyQueue) Push(x int)  {
	//输入的时候先检查输出栈是否为空
	//如果不为空那么将所有元素移动到输入栈
	//这时重新输入元素也就是后输入的元素，重新输入的元素将位于输出栈的底部最后输出
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		this.stack = append(this.stack,val)
	}
	this.stack = append(this.stack,x)
}


func (this *MyQueue) Pop() int {
	//如果输入栈不为空，那么将输入栈中的所有元素放入输出栈中
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back,val)
	}
	//如果输出栈==0直接返回
	if len(this.back) == 0 {
		return 0
	}
	//返回输出栈的最后元素
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}


func (this *MyQueue) Peek() int {
	//如果输入栈不为空，那么将输入栈中的所有元素放入输出栈中
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back,val)
	}
	//如果输出栈==0直接返回
	if len(this.back) == 0 {
		return 0
	}
	//返回输出栈的最后元素
	val := this.back[len(this.back)-1]
	return val
}

//直接判断两个栈中的元素是否为空
func (this *MyQueue) Empty() bool {
	return len(this.back)==0 && len(this.stack)==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
