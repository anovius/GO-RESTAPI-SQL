package main

import (
	"fmt"
	"restapi/model"
)

func main() {

	todo := model.Todo{
		ID:     1,
		Body:   "Hello World",
		Status: "Done",
	}

	fmt.Println(todo.Body)
}
