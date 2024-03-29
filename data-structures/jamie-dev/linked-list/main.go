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

func (l *LinkedList) prepend(n *Node) {
	temp := l.head
	l.head = n
	l.head.next = temp
	l.length++
}

func (l LinkedList) printList() {
	temp := l.head
	if temp == nil {
		fmt.Print("empty")
	} else {
		for temp != nil {
			fmt.Print(temp.data)
			temp = temp.next
			if temp != nil {
				fmt.Print(" -> ")
			}
		}
	}
	fmt.Print("\n")
}

func main() {
	aList := LinkedList{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 4}
	node5 := &Node{data: 5}
	node6 := &Node{data: 6}

	aList.prepend(node1)
	aList.prepend(node2)
	aList.prepend(node3)
	aList.prepend(node4)
	aList.prepend(node5)
	aList.prepend(node6)

	aList.printList()
}
