package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//主体思路是双指针,使fast指针比slow指针快n+1个节点，也就是间隔为n个节点
	//这样的话删除节点时会比较方便
	//那么当fast指针指向空的时候，slow.Next所指向要删除的指针
	//当tmp大于n也就是为n+1时两个指针开始同时移动
	//为什么不将fast与slow都设置为head同时将tmp设置为0呢
	//因为存在链表中只有一个节点，需要删除倒数第一个节点的情况
	//这时候就会删除失败
	newhead := &ListNode{}
	newhead.Next = head
	fast,slow := head,newhead
	tmp := 1  //tmp表示fast指针现在是第几个节点
	for fast != nil{
		fast = fast.Next
		if tmp > n {
			slow = slow.Next
		}
		tmp++
	}
	slow.Next =slow.Next.Next
	return newhead.Next
}