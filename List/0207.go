package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//对于本题的思路主要是如果存在交点
	//如果我们使p1p2指针分别同时遍历完headA和headB后没有存在相交情况
	//使得p1=headB p2 = headA 如果存在相交的情况那么必然会存在 p1 = p2
	//如果交换头指针后仍然有一个指针遍历到nil了那么就不存在交点
	p1, p2 := headA,headB
	len1,len2 := 0,0 //len1,len2分别存储两个链表的长度
	for p1 != nil {
		len1++
		p1 = p1.Next
	}
	for p2 != nil{
		len2++
		p2 = p2.Next
	}
	//如果有一个链表为nil那么直接返回nil
	if len1 == 0 || len2 == 0{
		return nil
	}
	tmp := len1 + len2//tmp等于总长
	p1, p2 = headA,headB
	for tmp >= 0{//因为存在一次重新指向head的操作所以不能是 != 0
		tmp--
		if p1 == p2{//如果此时节点相交那么直接返回p1或者p2
			return p1
		}
		//为什么要使用ifelse呢是因为每一次只能进行一次next操作
		if p1 == nil{
			p1 = headB
		}else{
			p1 = p1.Next
		}
		if p2 == nil{
			p2 = headA
		}else{
			p2 = p2.Next
		}
	}
	return nil
}