package main
//1.先找中点
//2.翻转后半部分
//3.合并链表
func reorderList(head *ListNode)  {
	midNode := findMidNode(head)
	tmp := midNode.Next
	midNode.Next = nil
	newHead := reverseList(tmp)
	mergeList(head,newHead)
}
func mergeList(l1,l2 *ListNode) {
	var l1tmp,l2tmp *ListNode
	for l1 != nil && l2 != nil {
		l1tmp = l1.Next
		l2tmp = l2.Next
		l1.Next = l2
		l1 = l1tmp
		l2.Next = l1
		l2 = l2tmp
	}
}
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next=newHead
		newHead = cur
		cur = tmp
	}
	return newHead
}
func findMidNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
