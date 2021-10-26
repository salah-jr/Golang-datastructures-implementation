package main

import "fmt"

const ArraySize = 7

type HashTable struct{
	array [ArraySize]*bucket
}

type bucket struct{
	head *bucketNode
}

type bucketNode struct{
	key string
	next *bucketNode
}

func (h *HashTable) Insert(key string){
	index := hash(key)	
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool{
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) bool{
	index := hash(key)
	return h.array[index].delete(key)
}

func (b *bucket) insert(key string){
	if !b.search(key){
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Key already exists!")
	}
}

func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil{
		if currentNode.key == key{
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(key string) bool{
	if b.head.key == key {
		b.head = b.head.next
		return true
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
			return true
		}
		previousNode = previousNode.next
	}
	return false
}

func hash(key string) int {
	sum := 0 
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func Init() *HashTable{
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main(){

	hashTable := Init()

	list := []string{
		"Randy",
		"Sarah",
		"John",
		"Alex",
		"Rick",
		"Sandy",
		"Pablo",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println(hashTable)

	//found
	fmt.Println(hashTable.Search("Rick"))
	// not found
	fmt.Println(hashTable.Search("Madona"))

	// Try to delete a node
	fmt.Println(hashTable.Delete("Sandy"))
	fmt.Println(hashTable.Search("Sandy"))

	// Try to delete the head
	fmt.Println(hashTable.Delete("Randy"))
	fmt.Println(hashTable.Search("Randy"))

	// Try to delete node that not exist
	fmt.Println(hashTable.Delete("Will"))

}