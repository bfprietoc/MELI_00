package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"meli_00/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mutant", controller.Mutant).Methods("POST")
	router.HandleFunc("/stats", controller.Stats).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
