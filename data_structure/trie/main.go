// This code is based on the following tutorial:
// https://www.youtube.com/watch?v=H-6-8_p88r0

package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{
		root: &Node{},
	}

	return result
}

func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, char := range w {
		charIndex := char - 'a'
		if currentNode.children[charIndex] == nil {
			// It does not exist
			// then we insert it
			currentNode.children[charIndex] = &Node{}
		}
		// It exists in the current node array
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, char := range w {
		charIndex := char - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func main() {
	trie := InitTrie()
	trie.Insert("supercalifragilisticexpialidocious")
	fmt.Println(trie.Search("supercalifragilisticexpialidocious"))
	fmt.Println(trie.Search("normalcalifragilisticexpialidocious"))
}
