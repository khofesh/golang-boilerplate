package main

import "fmt"

func main() {
	fmt.Println("Sets")
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
