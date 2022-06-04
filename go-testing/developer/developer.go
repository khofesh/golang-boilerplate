package developer

import "fmt"

// Developer
type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	var uniques []string
	check := make(map[string]int)

	for _, developer := range developers {
		check[developer.Name] = 1
	}

	fmt.Print(check)

	for name := range check {
		uniques = append(uniques, name)
	}

	return uniques
}
