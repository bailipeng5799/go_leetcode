package main
func reverseList(head *ListNode) *ListNode {
	var newhead *ListNode
	for head != nil{
		//tmp作为临时变量保存原来head节点后面的节点
		//交替替换newhead和head
		tmp := head.Next
		head.Next = newhead
		newhead = head
		head = tmp
	}
	return newhead
}
