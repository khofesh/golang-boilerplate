package main

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
}

// Node
// represents a link in our linked list
type Node struct {
	Value string
	Next  *Node
}

// Insert
// adds a new element to the start of our linked list
func (l *LinkedList) Insert(elem string) {
	node := Node{
		Value: elem,
		Next:  l.Head,
	}

	l.Head = &node
	l.Size++
}

// DeleteFirst
// removes the first element from our linked list
func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

// List
// iterates through all of the elements in our linked list and prints
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

// Search
// returns the first element that matches the string
// otherwhise it returns nil
func (l *LinkedList) Search(elem string) *Node {
	current := l.Head
	for current != nil {
		if current.Value == elem {
			return current
		}
		current = current.Next
	}

	return nil
}

// Delete
// removes an element from the linked list if it matches
// the value
func (l *LinkedList) Delete(elem string) {
	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Value == elem {
			previous.Next = current.Next
			l.Size--
		}
		previous = current
		current = current.Next
	}
}
