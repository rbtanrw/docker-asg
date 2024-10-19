package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to our webpage. Check /secret.")
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is not a secret.")
}

func handleRequest() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/secret", secretHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error Running Web Service")
	}
}

func main() {
	fmt.Println("Listening to port 8080...")
	handleRequest()
}
