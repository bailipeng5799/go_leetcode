package main

import "container/list"

type CQueue struct{
	stack1,stack2 *list.List//一个双向链表
}
func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),//创建一个双向链表返回头指针
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int)  {
	this.stack1.PushBack(value)//PushBack将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。
}

func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))//Back返回链表1的最后一个元素
			//remove删除链表1的最后一个元素，并且返回这个元素的val
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}
