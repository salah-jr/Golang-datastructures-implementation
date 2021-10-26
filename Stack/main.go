package main

import "fmt"

type Stack struct{
	items[] int
}

func (s *Stack) push(item int){
	s.items = append(s.items, item)
}
func (s *Stack) pop() int{
	index := len(s.items)-1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func main (){
	myStack := Stack{}

	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	fmt.Println("The stack after push: ",myStack)
	
	item := myStack.pop()
	fmt.Println("The item removed from stack is: ",item)
	fmt.Println("The stack after pop: ", myStack)
}
