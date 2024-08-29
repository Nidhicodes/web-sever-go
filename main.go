package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		formFilePath := filepath.Join(".", "static", "form.html")
		http.ServeFile(w, r, formFilePath)

	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "POST request successful")
		fmt.Fprintf(w, "Name: %s", name)
		fmt.Fprintf(w, "Address: %s", address)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello nidhi!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
