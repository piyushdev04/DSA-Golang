package main

import "fmt"

// Node represents an element in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the Linked List
type LinkedList struct {
	head *Node
}

// InsertAtEnd inserts a new node with given data at the end of the list
func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Delete deletes the first node with the given data
func (list *LinkedList) Delete(data int) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

// Print prints the linked list
func (list *LinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}

	// Insert elements
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)

	fmt.Println("Linked List:")
	list.Print()

	// Delete an element
	list.Delete(20)

	fmt.Println("Linked List after deleting 20:")
	list.Print()
}
