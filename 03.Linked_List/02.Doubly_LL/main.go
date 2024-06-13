package main

import "fmt"

// DNode represents an element in the doubly linked list
type DNode struct {
	data int
	next *DNode
	prev *DNode
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	head *DNode
}

// InsertAtEnd inserts a new node with given data at the end of the list
func (list *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &DNode{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
}

// Delete deletes the first node with the given data
func (list *DoublyLinkedList) Delete(data int) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next
		if list.head != nil {
			list.head.prev = nil
		}
		return
	}

	current := list.head
	for current != nil && current.data != data {
		current = current.next
	}

	if current != nil {
		if current.next != nil {
			current.next.prev = current.prev
		}
	}
}

// PrintForward prints the linked list from head to tail
func (list *DoublyLinkedList) PrintForward() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// PrintBackward prints the linked list from tail to head
func (list *DoublyLinkedList) PrintBackward() {
	current := list.head
	if current == nil {
		fmt.Println("nil")
		return
	}

	// Move to the end of the list
	for current.next != nil {
		current = current.next
	}

	// Print backwards
	for current.next != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}

func main() {
	list := &DoublyLinkedList{}

	// Insert elements
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)

	fmt.Println("Linked List Forward:")
	list.PrintForward()

	fmt.Println("Linked List Backward:")
	list.PrintBackward()

	// Delete an element
	list.Delete(20)

	fmt.Println("Linked List Forward after deleting 20:")
	list.PrintForward()

	fmt.Println("Linked List Backward after deleting 20:")
	list.PrintBackward()
}
