package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		version := os.Getenv("VERSION")
		if version == "" {
			version = "1.0.0"
		}
		fmt.Fprintf(w, "Hello from Docker and go! Version: %s", version)
	})
	log.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
