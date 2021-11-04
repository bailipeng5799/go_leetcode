package main
//k个一组翻转链表
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup( head *ListNode ,  k int ) *ListNode {
	if head == nil && k <= 1 {
		return head
	}
	lenHead := length(head) //获取链表的长度
	res := &ListNode{} //要返回的结果
	cur := res //每一次翻转之后的头节点的前一个指针节点
	tmpcnt := lenHead/k
	for i := 0; i < tmpcnt; i++ {
		var newHead *ListNode //这里不能定义为newHead := &ListNode{} 如果是这样的话那么每一次就会多出来一个0元素
		for j := 0; j < k; j++ { //链表翻转 头插
			tmp := head.Next
			head.Next = newHead
			newHead = head
			head = tmp
		}
		cur.Next = newHead
		for cur.Next != nil {
			cur = cur.Next
		}
	}
	cur.Next = head // 将最后剩下的小于k个的元素进行顺序添加
	return res.Next
}
func length(head *ListNode) int{
	res := 0
	for head != nil {
		res++
		head = head.Next
	}
	return res
}
