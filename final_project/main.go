package main

import (
	"final_project/datasource"
	"fmt"
)

func main() {
	_, err := datasource.New()
	if err != nil {
		fmt.Println("error start database", err)
		return
	}
}
