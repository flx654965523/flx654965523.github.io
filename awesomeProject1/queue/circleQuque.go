package queue

import (
	"errors"
	"fmt"
)

type circle struct {
	array [5]int
	tail int
	head int
	maxSize int
}

//入队列
func (this *circle)push(val int)error{
	err := this.isFull()
	if err != nil{
		this.array[this.tail] = val
		this.tail = (this.tail + 1) % this.maxSize
	}
	return nil
}

//出队列
func (this *circle)pop() (val int,err error) {
if this.isEmpty(){
return 0,errors.New("空队列")
}
val = this.head
this.head = (this.head + 1) % this.maxSize
return val,nil
}

//判断队列是否已满
func (this *circle)isFull() error{
	if (this.tail + 1) % this.maxSize == this.head{
		return errors.New("队列已满")
	}
	return nil
}

//判断队列是否为空
func (this *circle)isEmpty() bool{
	if this.tail == this.head{
		return true
	}
	return false
}

//判断队列中有多少元素
func (this *circle) size()  int{
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

//遍历
func (this *circle) list() error {
	if this.size() == 0{
		return errors.New("空队列")
	}
	temp := this.head
	for i:=0;i<this.size();i++{
		fmt.Println(temp,this.array[temp])
		temp = (temp +1) % this.maxSize
	}
	return nil
}


