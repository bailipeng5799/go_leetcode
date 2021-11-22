package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		Next : head,
	}
	pre := newHead
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			break
		}
		pre = head
		head = head.Next
	}

	return newHead.Next
}
