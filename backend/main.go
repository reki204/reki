package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func main() {
	http.HandleFunc("/", apiHandler)
	fmt.Println("Golang Server Start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
