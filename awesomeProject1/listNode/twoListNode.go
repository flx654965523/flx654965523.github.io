package listNode

import "log"

type TwoListNode struct {
	Pre *TwoListNode
	Next *TwoListNode
	Val int
}

//双向链表末尾添加元素
func AddNode(head *TwoListNode,newListNode *TwoListNode)  {
	temp := head

	for {
		if temp.Next == nil{
			temp.Next = newListNode
			newListNode.Pre = temp
			break
		}
		temp = temp.Next
	}
}

//双向链表中间插入元素
func InsertListNode(head *TwoListNode,newListNode *TwoListNode)  {
	temp := head

	for {
		if temp.Next == nil{
			break
		}
		if temp.Next.Val > newListNode.Val{
			newListNode.Next = temp.Next
			newListNode.Pre = temp
			temp.Next.Pre = newListNode
			temp.Next = newListNode
		}
		temp = temp.Next
	}
}

//向双向链表中删除元素
func delListNode(head *TwoListNode,val int)  {
	temp := head

	if temp.Next == nil{
		log.Fatalln("空链表")
		return
	}

	for{
		if temp.Next == nil{
			log.Fatalln("删除的值不存在")
			break
		}
		if temp.Next.Val == val{
			temp.Next = temp.Next.Next
			if temp.Next != nil{
				temp.Next.Pre = temp
			}
			break
		}
		temp = temp.Next
	}
}

//遍历双向链表
