package main

import (
	"log"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/html")
	w.Write([]byte("Hello world"))
}

func main() {
	http.HandleFunc("/", httpHandler)

	log.Println("Server started on :9000")
	log.Println("Visit http://localhost:9000")
	log.Println("Press Ctrl+C to stop the server")

	http.ListenAndServe(":9000", nil)
}
