package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (list *LinkedList) prepend(n *Node) {
	second := list.head
	list.head = n
	list.head.next = second
	list.length++
}

func (list *LinkedList) deleteNode(value int) {
	if list.length == 0 {
		fmt.Println("List Is Empty!")
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		list.length--
		return
	}
	node := list.head
	for node.next.data != value {
		if node.next.next == nil {
			fmt.Println("The node isn't found")
			return
		}
		node = node.next
	}
	node.next = node.next.next
	list.length--
}

func (list LinkedList) printListData() {
	valueToPrint := list.head
	for list.length != 0 {
		fmt.Printf("%d ", valueToPrint.data)
		valueToPrint = valueToPrint.next
		list.length--
	}
	fmt.Printf("\n")
}

func main() {
	myList := LinkedList{}
	node2 := Node{data: 65}
	node1 := Node{data: 28}
	node3 := Node{data: 44}
	node4 := Node{data: 93}
	node5 := Node{data: 15}
	node6 := Node{data: 57}

	myList.prepend(&node1)
	myList.prepend(&node2)
	myList.prepend(&node3)
	myList.prepend(&node4)
	myList.prepend(&node5)
	myList.prepend(&node6)

	//Try to delete node that doesn't exist in the list
	// myList.deleteNode(30)

	//Try to delete the head nodde
	// myList.deleteNode(57)

	//Try to delete node that ixusts in the list
	myList.deleteNode(44)

	fmt.Print("\n", "Here is the list.", "\n")
	myList.printListData()
}
