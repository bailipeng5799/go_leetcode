package main
//type ListNode struct {
//     Val int
//     Next *ListNode
//}
func getIntersectionNode(headA, headB *ListNode) *ListNode{
	//先让p遍历一遍链表A然后再去遍历链表B
	//再让q遍历一遍链表B然后再去遍历链表A如果两个结点相等那么就是第一个公共结点
	//如果都遍历结束同时为空那么就返回null
	if headA == nil || headB == nil{
		return nil
	}
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		}else{
			p = p.Next
		}
		if q == nil{
			q = headA
		}else {
			q = q.Next
		}
	}
	return p
}
