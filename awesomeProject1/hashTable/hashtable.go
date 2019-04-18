package hashTable

type User struct {
	name string
	no int32
	next *User
}

type ListNodes struct {
	Head *User
}

type HashTables struct {
	ListNode [7]ListNodes
}

func (hash *HashTables) Add(user *User)  {
	var index =  user.no % 7
	hash.ListNode[index].Add(user)
}

func (node *ListNodes)Add(user *User){
	temp := node.Head
	var pre *User = nil

	if temp == nil{
		node.Head = user
		return
	}

	for {
		if temp != nil{
			if temp.no > user.no{
				break
			}
			pre = temp
			temp = temp.next
		}else{
			break
		}
	}

	if temp == nil{
		pre.next = user
		user.next = temp
	}

}