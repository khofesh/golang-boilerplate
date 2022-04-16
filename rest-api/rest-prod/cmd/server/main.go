package main

import (
	"context"
	"fmt"
	"rest-prod/internal/db"
)

/*
  Run
  instantiate and start up our go app
*/
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected and pinged the database")

	return nil
}

func main() {
	fmt.Println("rest api")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
