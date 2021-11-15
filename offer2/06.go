package main
type ListNode struct {
    Val int
    Next *ListNode
}
func reversePrint(head *ListNode) []int {
	res := []int{}
	if head == nil {
		return res
	}
	for head != nil {
		res = append(res,head.Val)
		head = head.Next
	}
	for i := 0; i < len(res)/2; i++{
		res[i],res[len(res)-i-1]=res[len(res)-i-1],res[i]
	}
	return res
}