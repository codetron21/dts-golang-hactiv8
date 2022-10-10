package main

import (
	"assignment_2/database"
	"fmt"
)

func main() {
	_, err := database.Start()
	if err != nil {
		fmt.Println("error start database", err)
		return
	}
}
