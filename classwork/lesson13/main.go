package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method)
	// fmt.Println(r.Pattern)
	// fmt.Println(r.Host)
	// fmt.Println(r.RemoteAddr)

	type User struct {
		Name  string
		Email string
	}

	newUser := User{
		Name:  "Тёма",
		Email: "mrcat@example.com",
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	timeNow := time.Now().Format("02/01/2006 15:04")
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, newUser); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, timeNow); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//w.Write([]byte("Hello from index"))
	//http.ServeFile(w, r, "./templates/index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello drom about page")
}
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts page")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contacts", contactsHandler)

	log.Println("Server starting...")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println(err)
	}
}
