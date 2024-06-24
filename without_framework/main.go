package main

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"without.framework/router"
)

func main() {
	fmt.Println("Server listening on port 8080...")
	router := router.New()
	router.Get("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello, World!")
	})
	router.Post("/hello", func(res http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("error reading body: %s\n", err)
		}
		fmt.Fprintf(res, "Hello, %s!", body)
	})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
	http.Serve(l, router)
}
