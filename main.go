package main

import (
	"fmt"
	"net/http"
	"restapi/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetAll).Methods("GET")
	router.HandleFunc("/{id}", controller.GetOne).Methods("GET")
	router.HandleFunc("/create", controller.Create).Methods("POST")
	router.HandleFunc("/status", controller.ChangeStatus).Methods("PUT")
	router.HandleFunc("/delete/{id}", controller.Delete).Methods("DELETE")
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", router)
}
