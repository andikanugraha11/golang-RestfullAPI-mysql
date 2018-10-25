package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", returnAllUsers).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
