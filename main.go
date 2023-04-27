package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func fomrHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error in parsing form: %v", err)
		return
	}

	fmt.Fprintf(w, "Successfully submitted!")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Hi %v from %v", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", fomrHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
