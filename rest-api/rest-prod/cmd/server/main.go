package main

import (
	"context"
	"fmt"
	"rest-prod/internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and pinged the database")

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(context.Background(), "1"))

	return nil
}

func main() {
	fmt.Println("rest api")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
