package main
//206. 反转链表
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
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //如果输入的head为nil直接返回
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}