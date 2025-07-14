package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Everything changes!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Go API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

