package main

//203. 移除链表元素
// Definition for singly-linked list.
  type ListNode struct {
      Val int
      Next *ListNode
 }

func removeElements(head *ListNode, val int) *ListNode {
	//因为如果头节点要删除和别的节点删除元素的方法不同
	//所以我们自己构造一个头节点，这样所有节点的删除方法就会相同
	//使用cur作为中间变量用来循环
	//最后返回我们自己构建的头节点的Next
	l1 := &ListNode{}
	l1.Next = head
	cur := l1
	for cur != nil && cur.Next != nil{
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return l1.Next
}