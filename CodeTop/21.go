package main
//合并两个有序链表

//Definition for singly-linked list.
 type ListNode struct {
    Val int
     Next *ListNode
 }
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//对参数进行判断
	if l1 == nil && l2 == nil{
		return nil
	}else if l1 != nil && l2 == nil{
		return l1
	}else if l1 == nil && l2 != nil{
		return l2
	}
	//初始化一个头节点
	tmp := l1
	if l1.Val <= l2.Val{
		tmp = l1
		l1 = l1.Next
	}else{
		tmp = l2
		l2 = l2.Next
	}
	head := tmp//保存新链表的头指针
	//循环结束分为三种情况
	//1.两者恰好同时为空
	//2.l1为空且l2不为空，将l2拼接到新的链表后
	//3.l1不为空l2为空，将l1拼接到新的链表后
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil{
			if l1.Val <= l2.Val{
				tmp.Next = l1
				l1 = l1.Next
				tmp = tmp.Next
			}else{
				tmp.Next = l2
				l2 = l2.Next
				tmp = tmp.Next
			}
		}else if l1 != nil && l2 == nil{
			tmp.Next = l1
			break
		}else if l2 != nil && l1 == nil{
			tmp.Next = l2
			break
		}
	}
	return head
}