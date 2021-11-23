package main
//剑指 Offer 22. 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	resNode := head
	for head != nil {
		if k > 0 {
			k--
			head = head.Next
			continue
		}
		head = head.Next
		resNode = resNode.Next
	}
	return resNode
}
