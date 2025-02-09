package main

import (
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/html")
	w.Write([]byte("Hello world"))
}

func main() {
	http.HandleFunc("/", httpHandler)

	http.ListenAndServe(":9000", nil)
}
