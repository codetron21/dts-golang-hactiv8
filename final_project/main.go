package main

import (
	"final_project/controller"
	"final_project/datasource"
	"final_project/repository"
	"final_project/route"
	"final_project/service"
	"fmt"
)

func main() {
	db, err := datasource.Start()
	if err != nil {
		fmt.Println("error start database", err)
		return
	}

	repository := repository.New(db)

	service := service.New(&repository)

	ctl := controller.New(&service)

	err = route.StartServer(ctl)
	if err != nil {
		fmt.Println("error start server", err)
		return
	}
}
