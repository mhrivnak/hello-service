package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s!\n", hostname)
	})

	http.ListenAndServe(":8080", nil)
}
