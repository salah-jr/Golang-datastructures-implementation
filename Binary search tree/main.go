package main

import "fmt"

var count int

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (node *Node) insert(k int) {
	if k > node.key {
		if node.right == nil {
			node.right = &Node{key: k}
		} else {
			node.right.insert(k)
		}
	} else if k < node.key {
		if node.left == nil {
			node.left = &Node{key: k}
		} else {
			node.left.insert(k)
		}
	} else {
		fmt.Println("The key is exist in the tree")
	}
}

func (node *Node) search(k int) bool {
	count++
	if node == nil {
		return false
	}
	if k > node.key {
		return node.right.search(k)
	} else if k < node.key {
		return node.left.search(k)
	}
	return true
}
func main() {
	tree := Node{key: 100}

	tree.insert(50)
	tree.insert(300)
	tree.insert(40)
	tree.insert(84)
	tree.insert(20)
	tree.insert(89)
	tree.insert(206)
	tree.insert(378)
	tree.insert(79)
	tree.insert(67)
	tree.insert(1)

	fmt.Println(tree.search(84)) // exist
	fmt.Println(count)
	fmt.Println(tree.search(33)) // not exist
	fmt.Println(count)
	tree.insert(50) // key already in the tree

}
