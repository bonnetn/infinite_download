package main

import (
	"net/http"
	"log"
)

var DATA [1024 * 1024]byte

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("New connection!")
	w.Header().Add("Content-Type", "application/octet-stream")
	for {
		w.Write(DATA[:])
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
