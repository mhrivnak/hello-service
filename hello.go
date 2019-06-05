package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "Hello"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s from %s!\n", greeting, hostname)
	})

	http.ListenAndServe(":8080", nil)
}
