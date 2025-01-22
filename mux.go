package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome to the Home page!")
	if err != nil {
		fmt.Println(err)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "This is the About page.")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/about", handleAbout)

	err := http.ListenAndServe(":8089", mux)
	if err != nil {
		fmt.Println(err)
	}
}
