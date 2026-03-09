package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
)

type osData struct {
	Hostname     string
	OS           string
	Architecture string
	CPUs         int
	GoVersion    string
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	var data = osData{
		Hostname:     hostname,
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		CPUs:         runtime.NumCPU(),
		GoVersion:    runtime.Version(),
	}
	// Parse the template file
	tmpl := template.Must(template.ParseFiles("layout.html"))

	// Execute the template, writing the output to the http.ResponseWriter
	tmpl.Execute(w, data)

	w.Header().Set("content-type", "text/html")
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", httpHandler)

	log.Println("Server started on :" + port)
	log.Println("Visit http://localhost:" + port)
	log.Println("Press Ctrl+C to stop the server")

	http.ListenAndServe(":"+port, nil)
}
