package main
//对这道题存在疑问？？？
type MyLinkedList struct {
	//这个结构体用来初始化链表
	dummy *Node
}
type Node struct{
	Val int
	Next *Node
	Pre *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	//创建一个环形双端链表
	cur := &Node{
		Val : -1,
		Next : nil,
		Pre : nil,
	}
	cur.Next = cur
	cur.Pre = cur
	return MyLinkedList{cur}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	//当head == this.dummy时说明已经遍历一遍链表了
	for head  != this.dummy && index > 0{
		index--
		head = head.Next
	}
	if 0 != index{
		return -1
	}
	return head.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	//先申请一个Node结构体来存放元素
	tmp := &Node{
		Val : val,
		Next : this.dummy.Next,
		Pre : this.dummy,
	}
	//顺序不能相反，必须先更改之前第二个节点的前置指针
	this.dummy.Next.Pre = tmp//改变原来的头节点的pre指针
	this.dummy.Next = tmp
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	dummy := this.dummy // 拿到 头节点
	tmp := &Node{
		Val : val,
		Next : dummy, //末尾的下一个元素就是头节点
		Pre : dummy.Pre, // 末尾的前一个节点就是    之前头节点的上一个节点
	}
	//顺序不能反，必须先更改之前链表末尾元素的后置指针指向新的元素
	dummy.Pre.Next = tmp
	dummy.Pre = tmp//链表的头节点就是新插入在末尾的元素
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	head := this.dummy.Next
	for head != this.dummy && index > 0{
		index--
		head = head.Next
	}
	tmp := &Node{
		Val : val,
		Next : head,
		Pre : head.Pre,
	}
	//这两步不能反
	head.Pre.Next = tmp
	head.Pre  = tmp
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	head := this.dummy.Next
	for head.Next!= this.dummy && index > 0{
		index--
		head = head.Next
	}
	if index == 0 {
		head.Next.Pre = head.Pre
		head.Pre.Next = head.Next
	}
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */