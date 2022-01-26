// Backend Server

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", filter)
	http.ListenAndServe(":8090", nil)
}

func filter(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(body))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi there, I am the backend server for Order Browser %s", r.URL.Path[1:])
}
