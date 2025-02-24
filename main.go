package main

import (
	"log"
	"net/http"
)

type Time struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/time", timeHandler)

	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
