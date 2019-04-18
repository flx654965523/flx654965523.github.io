package main

import (
	"errors"
	"fmt"
)

type stack struct {
	maxSize int
	index int
	array [5]int
}

//入栈
func (this *stack)push(val int)(err error)  {
	if this.index == this.maxSize - 1{
		return errors.New("栈已满")
	}
	this.index++
	this.array[this.index] = val
	return nil
}

//遍历所有元素
func (this *stack)List()(err error){
	if this.index == -1{
		return errors.New("空栈")
	}
	for i:= this.index;i>-1;i--{
		fmt.Println(i,this.array[i])
	}
	return nil
}

//出栈
func (this *stack)pop()(val int,err error) {
	if this.index == -1{
		return 0,errors.New("空栈")
	}
	val = this.array[this.index]
	this.index--
	return val,nil
}

func main()  {
	stack := &stack{
		maxSize: 5,
		index:-1,
		array: [5]int{},
	}

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	stack.List()
	stack.push(6)
	stack.List()
	stack.pop()
	stack.List()
}