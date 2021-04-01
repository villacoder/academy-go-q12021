package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "I'm alive")
	})
	router.HandleFunc("/csv", getInfoCSV).Methods("GET")
	router.HandleFunc("/csv-by-id", getInfoCSVByID).Methods("GET")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
