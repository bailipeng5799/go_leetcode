package main
//两个链表的第一个公共结点
func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	// write code here
	p1, p2 := pHead1, pHead2  //p1 p2代表这个指针先从p1或者p2开始
	for p1 != p2 {
		if p1 == nil {
			p1 = pHead2
		}else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = pHead1
		}else {
			p2 = p2.Next
		}
	}
	return p2
}