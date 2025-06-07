package main

import (
	"log"
	"net/http"
	"os"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/html")
	w.Write([]byte("Hello world"))
}

func main() {
	var port = os.Getenv("PORT")
	http.HandleFunc("/", httpHandler)

	log.Println("Server started on :" + port)
	log.Println("Visit http://localhost:" + port)
	log.Println("Press Ctrl+C to stop the server")

	http.ListenAndServe(":"+port, nil)
}
