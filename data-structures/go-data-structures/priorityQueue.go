package main

import "errors"

// PriorityQueue
// queue data structure
type PriorityQueue struct {
	High []int
	Low  []int
}

// Enqueue
// add an element of type int to the end of the queue
func (q *PriorityQueue) Enqueue(elem int, highPriority bool) {
	if highPriority {
		q.High = append(q.High, elem)
	} else {
		q.Low = append(q.Low, elem)
	}
}

// Dequeue
// returns the first element from the queue
func (q *PriorityQueue) Dequeue() (int, error) {
	// if the length of the high priority != 0
	// return the first element from the high priority queue
	// otherwise if the length of the low priority != 0
	// return the first element from the low priority queue
	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}

	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low = q.Low[0], q.Low[1:]
		return firstElement, nil
	}

	return 0, errors.New("empty queue")
}

// Length
// returns the length of the queue
func (q *PriorityQueue) Length() int {
	return len(q.High) + len(q.Low)
}

// Peek
// returns the first element from the queue without updating queue
func (q *PriorityQueue) Peek() (int, error) {
	if (len(q.High)) != 0 {
		return q.High[0], nil
	}

	if len(q.Low) != 0 {
		return q.Low[0], nil
	}

	return 0, errors.New("empty queue")
}

// IsEmpty
// return true if queue is empty
func (q *PriorityQueue) IsEmpty() bool {
	return q.Length() == 0
}
