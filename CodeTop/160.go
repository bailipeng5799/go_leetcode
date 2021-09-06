package main
//相交链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//核心思想
//公共尾部为c
//headA距离相交点距离为a
//headB距离相交点距离为b
//那么设置tmpA遍历一次headA之后遍历headB
//tmpB先遍历headB然后遍历headA
//如果有相交点 tmpA = a + c + b  tmpB = b + c + a 此时二者相遇循环结束
//如果没有两者同时为空tmpA == tmpB循环结束
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpA,tmpB := headA,headB
	for tmpA != tmpB {
		if tmpA != nil{//如果等于nil说明headA遍历结束了
			tmpA = tmpA.Next
		}else{
			tmpA = headB
		}
		if tmpB != nil{
			tmpB = tmpB.Next
		}else{
			tmpB = headA
		}
	}
	return tmpA
}