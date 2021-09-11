package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	//创建一个虚拟头节点
	cur := &ListNode{
		Next : head,
	}
	pre := cur
	for head != nil && head.Next != nil{
		//让虚拟头节点的Next指向head.Next也就是第二个节点
		pre.Next = head.Next
		//tmp同来暂时保存第三个节点
		tmp := head.Next.Next
		//头节点的Next指向第三个节点
		head.Next = tmp
		//虚拟头节点所指向的第二个节点的Next指向原来的第一个节点也就是头节点
		pre.Next.Next = head
		//前两个节点翻转完成让原来第一个节点也就是现在的第二个节点指向tmp
		head.Next = tmp
		//虚拟头节点可以看为下一组需要翻转的节点的前一个节点也就是第二个节点
		pre = head
		//头节点就是需要翻转的两个节点中的第一个
		head = tmp
	}
	return cur.Next
}
