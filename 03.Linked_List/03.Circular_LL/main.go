package main

import "fmt"

// Node represents an element in circular linked list
type Node struct {
	data int
	next *Node
}

// CircularLinkedList represents the circular linked list
type CircularLinkedList struct {
	head *Node
}

// InsertAtEnd inserts a new node with given data at the end of the list
func (list *CircularLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
	} else {
		current := list.head
		for current.next != list.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = list.head
	}
}

// Delete deletes the first node with the given data
func (list *CircularLinkedList) Delete(data int) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		if list.head.next == list.head {
			list.head = nil
			return
		}

		current := list.head
		for current.next != list.head {
			current = current.next
		}
		current.next = list.head.next
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != list.head && current.next.data != data {
		current = current.next
	}

	if current.next.data == data {
		current.next = current.next.next
	}
}

// Print prints the circular linked list
func (list *CircularLinkedList) Print() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := list.head
	for {
		fmt.Printf("%d -> ", current.data)
		current = current.next
		if current == list.head {
			break
		}
	}
	fmt.Println("Head")
}

func main() {
	list := &CircularLinkedList{}

	// Insert elements
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)

	fmt.Println("Circular Linked List:")
	list.Print()

	// Delete an element
	list.Delete(20)

	fmt.Println("Circular Linked List after deleting 20:")
	list.Print()
}
