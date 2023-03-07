package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" && r.Method == "GET" {
		fmt.Fprintf(w, "Hello")
		return
	}
	http.Error(w, "404 not found", http.StatusNotFound)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v\n", err)
		return
	}
	fmt.Fprintf(w, "POST request Success\n")
	fmt.Fprintf(w, "==========================\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name : %s\n", name)
	fmt.Fprintf(w, "address : %s\n", address)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	port := ":8080"

	// route
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	// log connection
	fmt.Printf("Starting server at port : 8080")
	

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
	
}