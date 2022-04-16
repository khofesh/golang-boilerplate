package main

import "fmt"

/*
  Run
  instantiate and start up our go app
*/
func Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {
	fmt.Println("rest api")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
