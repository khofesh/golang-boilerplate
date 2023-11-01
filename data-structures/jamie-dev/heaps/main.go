package main

import "fmt"

type MaxHeap struct {
	theslice []int
}

// Insert adds an element to the heap
func (m *MaxHeap) Insert(key int) {
	m.theslice = append(m.theslice, key)

	m.maxHeapifyUp(len(m.theslice) - 1)
}

func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.theslice[parent(index)] < m.theslice[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MaxHeap) Extract() int {
	extracted := m.theslice[0]

	L := len(m.theslice) - 1

	if len(m.theslice) == 0 {
		fmt.Println("array length is 0 ")
		return -1
	}

	/**
	 * take out the last index and put it in the root
	 */
	m.theslice[0] = m.theslice[L]
	m.theslice = m.theslice[:L]

	m.maxHeapifyDown(0)

	return extracted
}

func (m *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(m.theslice) - 1
	L, r := left(index), right(index)

	childToCompare := 0

	// while index has at least one child
	for L <= lastIndex {
		if L == lastIndex { // when index is the only child
			childToCompare = L
		} else if m.theslice[L] > m.theslice[r] { // when left child is larger
			childToCompare = L
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child
		// and swap if smaller
		if m.theslice[index] < m.theslice[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			L, r = left(index), right(index)
		} else {
			return
		}
	}
}

// parent gets the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left gets the left child index
func left(i int) int {
	return 2*i + 1
}

// right gets the right child index
func right(i int) int {
	return 2*i + 2
}

func (m *MaxHeap) swap(i1, i2 int) {
	m.theslice[i1], m.theslice[i2] = m.theslice[i2], m.theslice[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	// m.Extract()
	// fmt.Println(m)

	for i := 0; i < len(buildHeap)-1; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
