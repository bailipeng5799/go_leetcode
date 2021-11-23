package main
//剑指 Offer 24. 反转链表
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}
	return newHead
}
//递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}else {
		newHead := reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return newHead
	}
}
