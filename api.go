package main

import (
"fmt"
"net/http"
"log"
)

var (
	counter int
	port    = 9000
)

func count(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	counter++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("{'Message': 'Hello DevFest Granada 2017'}"))
}

func stats(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{'Visits': %d}", counter)))
}

func main() {
	http.HandleFunc("/favicon.ico", nil)
	http.HandleFunc("/", count)
	http.HandleFunc("/stats", stats)
	log.Println("Listening at port ", port)
	log.Panic(
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
