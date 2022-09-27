package controller

import (
	"encoding/json"
	"net/http"
	"restapi/config"
	"restapi/model"
	"strconv"

	"github.com/gorilla/mux"
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

//Get one

func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	defer db.Close()

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var todo model.Todo

	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	err := row.Scan(&todo.ID, &todo.Body, &todo.Status)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(todo)
}

// Create Todo

func Create(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	stmt, err := db.Prepare("INSERT INTO todos(body) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(todo.Body)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
}

//change todo status

func ChangeStatus(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	stmt, err := db.Prepare("UPDATE todos SET status = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(todo.Status, todo.ID)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
}

//delete todo

func Delete(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	db := config.Connect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
}
