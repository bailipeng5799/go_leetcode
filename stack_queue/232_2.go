package main
type MyQueue struct {
	stack []int
	back []int
}

//返回一个初始化的MyQueue
func Constructor() MyQueue {
	return MyQueue{
		stack : make([]int,0),
		back : make([]int,0),
	}
}

//输入
//判断输出栈是否为空如果不为空将输出栈中的元素从栈顶到栈底按序压入输入栈中
//最后将将要添加的元素压入输入栈中
func (this *MyQueue) Push(x int)  {
	// for len(this.back) != 0 {
	//     val := this.back[len(this.back)-1]
	//     this.back = this.back[:len(this.back)-1]
	//     this.stack =  append(this.stack,val)
	// }
	// this.stack = append(this.stack,x)
	this.stack = append(this.stack,x)
}


func (this *MyQueue) Pop() int {
	if len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		return val
	}
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back,val)
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}


func (this *MyQueue) Peek() int {
	if len(this.back) != 0 {
		return this.back[len(this.back)-1]
	}
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back,val)
	}
	return this.back[len(this.back)-1]
}

//判断栈是否为空
func (this *MyQueue) Empty() bool {
	return len(this.back) == 0 && len(this.stack) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
