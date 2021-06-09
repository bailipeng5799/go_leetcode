package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//链表翻转头插法
//相当于把head链表上的结点一个一个拆下来使用头插法插入到p链表上
func reverseList(head *ListNode) *ListNode {
	var p *ListNode
	for head != nil{
		temp := head.Next
		head.Next = p
		p = head
		head = temp
	}
	return p
}
//就地反转
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	beg := head
	end := head.Next
	for end != nil{
		beg.Next = end.Next//将end结点从链表中取出来
		end.Next = head
		head = end
		end = beg.Next
	}
	return head
}
//递归翻转
//当遍历到最后一个结点时new_head将会不变
//对head和head.Next进行操作，可以看成倒数第二个结点和最后一个结点
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}else{
		new_head := reverseList3(head.Next)
		head.Next.Next = head
		head.Next = nil
		return new_head
	}
}