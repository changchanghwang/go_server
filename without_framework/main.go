package main

import (
	"fmt"
	"log"
	"net/http"

	accountHandler "without.framework/handlers/account"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			accountHandler.AddAccount(w, r)
		}
	})

	fmt.Println("Server running on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
