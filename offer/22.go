package main
//快慢指针中间间隔K
func getKthFromEnd(head *ListNode, k int) *ListNode {
	left := head
	right := head
	for right != nil{
		if k != 0{
			right=right.Next
			k--
			continue
		}
		left = left.Next
		right = right.Next
	}
	return left
}