package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	    Val int
	    Next *ListNode
	}
func reversePrint4(head *ListNode) []int {

}
func reversePrint(head *ListNode) []int {
	if head == nil {return nil}
	//先通过递归函数进入到最后一个结点的next结点也即是空，然后在此基础上给添加int类型的val
	//添加最后一个以此类推
	return append(reversePrint(head.Next), head.Val)
}
//核心思想创建一个新的链表使用头插法将题目链表按序进行插入 达到翻转链表的效果
func reversePrint2(head *ListNode) []int{
	if head == nil{
		return nil
	}
	var newHead *ListNode
	res := []int{}
	for  head != nil{
		tempnode := head.Next
		head.Next=newHead
		newHead=head
		head=tempnode
	}
	for newHead != nil{
		res=append(res,newHead.Val)
		newHead=newHead.Next
	}
	return res

}
//先顺序到数组里面然后在反转数组
func reversePrint3(head *ListNode) []int {
	if head == nil{
		return nil
	}
	res := []int{}
	for head!=nil{
		res=append(res,head.Val)
		head=head.Next
	}
	for i,j := 0,len(res)-1;i < j; i, j = i+1, j-1{
		temp:=res[i]
		res[i]=res[j]
		res[j]=temp
	}
	return res
}


