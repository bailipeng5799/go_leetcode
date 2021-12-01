package main

//Definition for singly-linked list.
type ListNode struct {
      Val int
      Next *ListNode
}

/*
 核心思想就是先得到整个链表长度，然后判断需要翻转多少组，最后将每一组翻转然后拼接，最后在加上没有达到k的元素*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	length := GetLength(head)
	res := &ListNode{}
	cur := res
	tmpcnt := length / k
	for i := 0; i < tmpcnt; i++ {
		var newHead *ListNode
		for j := 0; j < k; j++ {
			tmp := head.Next
			head.Next = newHead
			newHead = head
			head=tmp
		}
		cur.Next = newHead
		for cur.Next != nil {
			cur = cur.Next
		}
	}
	cur.Next = head
	return res.Next
}
func GetLength(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}
