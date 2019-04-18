package main

import (
	"awesomeProject1/listNode"
	"awesomeProject1/models"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)
//切片结构体排序
type students []models.Student
var  studentype students

func (s students)Len() int{
	return len(studentype)
}

func (s students)Less (i, j int) bool {
	return studentype[i].Score < studentype[j].Score
}

func (s students)Swap (i, j int) {
	studentype[i], studentype[j] =  studentype[j],studentype[i]
}
func main1(){
	for i:=0;i<10;i++{
		student := models.Student{
			Name: "学生"+strconv.Itoa(i),
			Age : i,
			Score: rand.Intn(100),
		}
		studentype = append(studentype, student)
	}

	fmt.Println(studentype)
	sort.Sort(studentype)
	fmt.Println(studentype)
}

//类型断言判断
func contirmtype(items... interface{}){
	for index,value := range items{
		switch value.(type) {
		case bool:
			fmt.Printf("第%v个参数的类型是bool",index,value)
		case int:
			fmt.Printf("第%v个参数的类型是int",index,value)
		case float32:
			fmt.Printf("第%v个参数的类型是float32",index,value)
		case models.Student:
			fmt.Printf("第%v个参数的类型是Student",index,value)
		case *models.Student:
			fmt.Printf("第%v个参数的类型是*Student",index,value)
		}
	}
}

func main2()  {
	var i int = 0
	var f float32 = 2.14
	var student models.Student
	var stu *models.Student

	contirmtype(i,f,student,stu)
}

//map实例,如果不存在用户名，给用户名加个昵称和密码，有则修改密码
func modifyPwd(user map[string]map[string]string,name string) {
	if user[name] != nil{
		user[name]["pwd"] = "888888"
	}else{
		user[name] = make(map[string]string,10)
		user[name]["nickname"] = "flx"
		user[name]["pwd"] = "888888"
	}
}

func main3(){
	users := make(map[string]map[string]string,10)
	users["fulixin"] = make(map[string]string,10)
	users["fulixin"]["flx"] = "666666"
	modifyPwd(users,"fulixin")
	fmt.Println(users)

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i:=0;i<len(nums);i++{
		s := target - nums[i]
		if _,ok := m[s];ok{
			return []int{m[s],i}
		}
		m[nums[i]] = i
	}
	return nil
}

func removeDuplicates(nums []int) int {
	var j = 0
	for i:=1;i<len(nums);i++{
		if nums[i] != nums[j]{
			nums[j] = nums[i]
			j++
		}
	}
	return j+1
}

var listnode = listNode.OneListNode{}

func main(){
	onelistnode1 := &listNode.OneListNode{
						Val:1,
					}

	onelistnode2 := &listNode.OneListNode{
						Val:2,
					}

	onelistnode3 := &listNode.OneListNode{
						Val:3,
					}
	onelistnode4 := &listNode.OneListNode{
					Val:4,
	}
	onelistnode5 := &listNode.OneListNode{
					Val:5,
	}

	head := &listNode.OneListNode{}
	listnode.InsertNodeFormLast(head,onelistnode1)
	listnode.InsertNodeFormLast(head,onelistnode2)
	listnode.InsertNodeFormLast(head,onelistnode3)
	listnode.InsertNodeFormLast(head,onelistnode5)
	listnode.InsertNodeFormVal(head,onelistnode4)
	listnode.DelListNode(head,4)
	listNode.
	listnode.GetAllListNode(head)
}




