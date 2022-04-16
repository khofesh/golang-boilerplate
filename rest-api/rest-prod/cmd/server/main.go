package main

import (
	"fmt"
	"rest-prod/internal/comment"
	"rest-prod/internal/db"
	transportHttp "rest-prod/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("rest api")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
