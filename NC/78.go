package main
type ListNode struct{
	  Val int
	  Next *ListNode
}
//NC78 反转链表
func ReverseList1(pHead *ListNode) *ListNode{
	var newHewad *ListNode
	for pHead != nil {
		tmp := pHead.Next
		pHead.Next = newHewad
		newHewad = pHead
		pHead = tmp
	}
	return newHewad
}



//超时了
func ReverseList( pHead *ListNode ) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	newHead := ReverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead = nil
	return newHead
}
