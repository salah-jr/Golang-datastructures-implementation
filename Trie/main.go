package main

import "fmt"

// Number of possible characters in the trie
const AlphabetSize = 26

// Represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Represents the trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// Initialize new trie
func initTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will add a word to the trie
func (t *Trie) insert(word string) {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take a word and return true if exists in the trie
func (t *Trie) search(word string) bool {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := initTrie()

	myTrie.insert("alex")
	myTrie.insert("john")
	myTrie.insert("jack")
	myTrie.insert("sarah")
	myTrie.insert("lynda")

	// found : will return true
	fmt.Println(myTrie.search("sarah"))
	fmt.Println(myTrie.search("alex"))

	// not found : will return false
	fmt.Println(myTrie.search("doe"))
	fmt.Println(myTrie.search("mariam"))
}
