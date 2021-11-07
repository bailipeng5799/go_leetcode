package main
//双指针法 快指针比慢指针多走n
func removeNthFromEnd( head *ListNode ,  n int ) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	fast, slow, pre := head, head, head
	for n > 0{
		n--
		fast = fast.Next
	}
	if fast == nil { //如果fast == nil 说明此链表长度为n直接删除第一个元素
		return head.Next
	}
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
		if fast == nil {
			pre.Next = slow.Next
		}
	}
	return head
}
