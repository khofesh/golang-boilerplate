package main

import "fmt"

func setsSection() {
	fmt.Println("Sets Section")
	planets := NewSet()

	planets.Add("Earth")
	planets.Add("Venus")
	planets.Add("Mars")
	planets.Add("Earth")
	planets.Add("Pluto")

	planets.Delete("Pluto")

	planets.List()

	fmt.Println(planets.Contains("Mars"))
}

func queuesSection() {
	fmt.Println("Queues Section")

	queue := Queue{}
	fmt.Println(queue)
	queue.Enqueue(1)
	fmt.Println(queue)
	queue.Enqueue(2)
	fmt.Println(queue)

	elem, _ := queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)
}

func priorityQueueSection() {
	fmt.Println("Priority Queue")

	queue := PriorityQueue{}
	queue.Enqueue(1, true)
	fmt.Println(queue)
	queue.Enqueue(10, false)
	fmt.Println(queue)

	elem, _ := queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)

	queue.Enqueue(2, true)
	fmt.Println(queue)

	elem, _ = queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)
}

func stackSection() {
	fmt.Println("Stacks Section")

	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	peek1, _ := stack.Peek()
	fmt.Println(peek1)

	fmt.Println(stack.Length())

	elem1, _ := stack.Pop()
	fmt.Println(elem1)
	elem2, _ := stack.Pop()
	fmt.Println(elem2)
	elem3, _ := stack.Pop()
	fmt.Println(elem3)

	fmt.Println(stack.IsEmpty())
}

func main() {
	setsSection()

	queuesSection()

	priorityQueueSection()

	stackSection()
}
