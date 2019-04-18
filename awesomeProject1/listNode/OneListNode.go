package listNode

import (
	"fmt"
	"log"
)

//单链表
type OneListNode struct {
	Val int
	Next * OneListNode
}

//在链表的末尾添加
func(*OneListNode) InsertNodeFormLast(head *OneListNode, newListNode *OneListNode){
	temp := head

	for {
		if temp.Next == nil{
			break
		}
		temp = temp.Next
	}
	temp.Next = newListNode
}

//在链表的指定位置添加
func (*OneListNode)InsertNodeFormVal(head *OneListNode, newListNode * OneListNode)  {
	temp := head

	for {
		if temp.Next == nil{
			break
		}
		if temp.Next.Val > newListNode.Val{  //插入到temp后面
			break
		}
		temp = temp.Next
	}

	newListNode.Next = temp.Next
	temp.Next = newListNode

}

//获取链表所有元素
func (*OneListNode)GetAllListNode(head * OneListNode){
	temp := head

	if temp.Next == nil{
		log.Fatalln("没有任何数据")
		return
	}

	for {
		if temp.Next == nil{
			break
		}
		temp = temp.Next
		fmt.Print(temp.Val)
	}
}

//删除链表指定元素
func (*OneListNode)DelListNode(head *OneListNode,val int){
	temp := head

	if temp.Next == nil{
		log.Println("空链表")
		return
	}

	for {
		if temp.Next == nil{
			log.Println("没有指定要删除的元素")
			return
		}
		if temp.Next.Val == val{
			temp.Next = temp.Next.Next
			break
		}
		temp = temp.Next
	}
}

//修改链表指定元素
func(*OneListNode) ModifylistNode(head *OneListNode,val int,modifyVal int)  {
	temp := head
	if temp.Next == nil{
		log.Fatalln("空链表")
		return
	}
	for {
		if temp.Next == nil{
			log.Fatalln("不存在的修改值")
			return
		}

		if temp.Next.Val == val{
			temp.Next.Val = modifyVal
			break
		}
		temp = temp.Next
	}
}


