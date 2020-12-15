package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./login")))

	log.Fatal(http.ListenAndServe(":9091", nil))
}
