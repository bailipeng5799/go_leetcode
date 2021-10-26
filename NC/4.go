package main
func hasCycle( head *ListNode ) bool {
	// write code here
	fast,slow := head,head
	for fast != nil && fast.Next != nil { //如果判断fast.Next不为nil那么快指针遍历后最多为nil 不会进入循环
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
