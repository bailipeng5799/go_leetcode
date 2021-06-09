package main

type ListNode struct {
	Val int
    Next *ListNode
}
//创建一个新的链表进行操作
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p:=head
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			p.Next=l1
			l1=l1.Next
		}else{
			p.Next=l2
			l2=l2.Next
		}
		p=p.Next
	}
	if l1 != nil{
		p.Next = l1
	}
	if l2 != nil{
		p.Next = l2
	}
	return head.Next
}
//一直把val小的链表放在l1
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}
