package main

import (
	"fmt"
	"log"
	"net/http"
)

var listenAddr string = ":8000"

func main() {
	mux := http.NewServeMux()

	hello := func(res http.ResponseWriter, req *http.Request) {
		message := []byte("<h1 style=\"color: cyan\">Hello world from Go</h1>")
		res.Write(message)
	}

	mux.HandleFunc("/", hello)

	fmt.Printf("Server running on port %s\n", listenAddr[1:])
	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		log.Fatal(err)
	}
}
