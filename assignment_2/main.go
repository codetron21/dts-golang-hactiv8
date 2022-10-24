package main

import (
	"assignment_2/controller"
	"assignment_2/database"
	"assignment_2/routes"
	"fmt"
)

func main() {
	// setup db
	db, err := database.Start()
	if err != nil {
		fmt.Println("error start database", err)
		return
	}

	// controller
	ctl := controller.New(db)

	err = routes.StartServer(ctl)
	if err != nil {
		fmt.Println("error start server", err)
		return
	}
}
