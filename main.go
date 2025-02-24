package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Time struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	curentTime := Time{Time: time.Now().Format(time.RFC3339)}
	response, err := json.Marshal(curentTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)

    err := http.ListenAndServe(":8795", nil)
    if err != nil {
      log.Fatalf("Server failed to start: %v", err)
    }
}
