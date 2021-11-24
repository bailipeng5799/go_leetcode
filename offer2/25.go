package main
//递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	}else if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoLists(l1.Next,l2)
	return l1
}
//
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		}else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	}else {
		p.Next = l1
	}
	return head.Next
}