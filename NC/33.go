package main
func Merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	newHead := &ListNode{}
	res := newHead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			newHead.Next = pHead1
			pHead1 = pHead1.Next
		}else{
			newHead.Next = pHead2
			pHead2 = pHead2.Next
		}
		newHead = newHead.Next
	}
	if pHead1 == nil {
		newHead.Next = pHead2
	}
	if pHead2 == nil {
		newHead.Next = pHead1
	}
	return res.Next
}
