package main

import "fmt"

func TheArrays() {
	fmt.Println("Arrays")

	myArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(myArray)

	for i, v := range myArray {
		fmt.Printf("Index: %d\n", i)
		fmt.Println(v)
	}
}
