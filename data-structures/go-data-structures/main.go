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

func main() {
	setsSection()

	queuesSection()

	priorityQueueSection()
}
