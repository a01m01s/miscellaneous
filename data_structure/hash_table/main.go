// This code is based on the (https://www.youtube.com/watch?v=JCXLwfKMWOU) tutorial
// A hash table for strings

package main

import "fmt"

const ArraySize = 7

// Hash table structure
// which will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket structure
// which is a linked list
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

// Hash table methods
func hash(key string) int {
	// is the summation of the ASCII codes of the key elements
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket methods
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Already exists")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
	} else {

		previousNode := b.head
		for previousNode.next != nil {
			if previousNode.next.key == k {
				previousNode.next = previousNode.next.next
			}
			previousNode = previousNode.next
		}
	}
}

// Initializing the hash table
func Init() *HashTable {
	// Will create a bucket in each slot of the hash table
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"Eric",
		"Kenny",
		"Adele",
		"Emma",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}

}
