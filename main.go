package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" || r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	if _, err := fmt.Fprint(w, "Hello"); err != nil {
		http.Error(w, "Can't get page", http.StatusNotFound)
	}
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err :%v", err)
		return
	}
	if _, err := fmt.Fprintln(w, "POST request successful"); err != nil {
		http.Error(w, "POST request was not successful", http.StatusInternalServerError)
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	if _, err := fmt.Fprintf(w, "Name = %s\nAddress = %s", name, address); err != nil {
		http.Error(w, "Can't Get the Name and Address", http.StatusInternalServerError)
	}
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
