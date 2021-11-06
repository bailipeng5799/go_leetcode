package main
//本题使用快慢指针的算法
//fast指针每次走两步，slow指针每次走一步，如果存在环那么必然相交
//s代表起始位置到入环点的距离   b代表入环点到相遇点的距离  c 代表从相遇点到入环点的距离
//这时候fast = s + n(b+c)+b  slow = s + b    fast = 2 * slow
//          (n-1)b + nc = s = c+(n−1)(b+c)
// 从相遇点到入环点的距离加上 n-1 圈的环长，恰好等于从链表头部到入环点的距离。
func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	fast, slow := pHead, pHead
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			pre := pHead
			for pre != slow {
				pre = pre.Next
				slow = slow.Next
			}
			return pre
		}
	}
	return nil

}
