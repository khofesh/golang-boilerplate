package linting

import "fmt"

func checkFlag(flag bool) bool {
	if flag == true {
		return true
	} else {
		return false
	}
}

func errChecking() (int, error) {
	a := 1
	a = 2
	fmt.Println(a)

	return 0, nil
}
