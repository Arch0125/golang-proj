package main

import (
	"log"
	"net/http"
	"fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed);
		return;
	}

	if r.URL.Path != "/hello" {
		http.Error(w, "404 : Not found", http.StatusNotFound);
		return;
	}

	fmt.Fprintf(w, "Hello, World!");
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err);
		return;
	}

	fmt.Fprintf(w, "POST Success");
	name := r.FormValue("name");
	email := r.FormValue("email");
	fmt.Fprintf(w, "Name: %s", name);
	fmt.Fprintf(w, "Email: %s", email);

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/", fileServer);
	http.HandleFunc("/hello", helloHandler);
	http.HandleFunc("/form", formHandler);

	fmt.Println("Listening on port 8080");
	log.Fatal(http.ListenAndServe(":8080", nil));
}