package controller

import (
	"encoding/json"
	"net/http"
	"restapi/config"
	"restapi/model"
)

// Get All Todo's

func GetAll(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var todos []model.Todo

	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Body, &todo.Status)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
