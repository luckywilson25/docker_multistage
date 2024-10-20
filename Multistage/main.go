package main

import (
	"fmt"
	"net/http"
)

func appsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome back!")
}

func haiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi..")
}

func main() {
	http.HandleFunc("/apps", appsHandler)
	http.HandleFunc("/hai", haiHandler)

	fmt.Println("Server is running at http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error Server", err)
	}
}