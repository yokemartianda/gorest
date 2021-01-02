package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Running http server")
	log.Fatal(http.ListenAndServe(":8080", RestRouter()))
}