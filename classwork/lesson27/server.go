package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"status": "ok"}`))
}
func main() {
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Println(err)
		return
	}
}
