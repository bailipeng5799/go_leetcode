package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val{
		head=head.Next
		return head
	}
	p:=head
	for p.Next != nil{
		if p.Next.Val == val{
			p.Next=p.Next.Next
			return head
		}
		p=p.Next
	}
	return head
}