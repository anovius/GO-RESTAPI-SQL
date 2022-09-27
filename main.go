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
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", router)
}
