package main

import "fmt"

type Queue struct{
	items[] int
}

func (s *Queue) enqueue(item int){
	s.items = append(s.items, item)
}
func (s *Queue) dequeue() int{
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func main (){
	myQueue := Queue{}

	myQueue.enqueue(100)
	myQueue.enqueue(200)
	myQueue.enqueue(300)
	fmt.Println("The Queue after enqueue: ",myQueue)
	
	item := myQueue.dequeue()
	fmt.Println("The item removed from Queue is: ",item)
	fmt.Println("The Queue after dequeue: ", myQueue)
}
