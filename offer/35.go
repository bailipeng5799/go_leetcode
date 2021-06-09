package main

 //Definition for a Node.
type Node struct {
   Val int
   Next *Node
    Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil{
		return nil
	}
	//必须要把Random置为空
	//第二次循环map中的结点添加Random所指向的结点时需要判断条件
	//如果不置为空母链表的Random为null
	//那么复制后的子链表就会指向原来的链表中此结点Random指向的Node
	//没有进行完全复制
	newHead := Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	p := head.Next //作为代替head使用
	pre := &newHead//用来代替map中的value
	listMap := make(map[*Node]*Node)
	listMap[head] = pre
	for p != nil{
		tmp := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = tmp
		listMap[p] = tmp
		p = p.Next
		pre = pre.Next
	}
	//第二次遍历
	p = head
	twoNewP := &newHead//第二次进行遍历将random指针指向map中对应得key得value
	for p != nil{
		if p.Random != nil{
			twoNewP.Random = listMap[p.Random]//p.Random作为key读取到新链表结点得value
		}
		p=p.Next
		twoNewP = twoNewP.Next
	}
	return &newHead
}
