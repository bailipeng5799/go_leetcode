package main
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
	tmpcnt := length/k
	res := &ListNode{}
	cur := res
	for i := 0; i < tmpcnt; i++ {
		var newHead *ListNode
		for j := 0; j < k; j++ {
			tmp := head.Next
			head.Next = newHead
			newHead = head
			head = tmp
		}
		cur.Next = newHead
		for cur.Next != nil { // 找到翻转后需要链接的前一个节点
			cur = cur.Next
		}
	}
	cur.Next = head //将剩下的不足k个元素进行添加
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
