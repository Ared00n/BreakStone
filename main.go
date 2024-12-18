package main

import (
	"log"
	"net/http"
)

func main() {
	DB()
	defer CloseDB()

	links := setuplinks()

	if err := http.ListenAndServe(":8080", links); err != nil {
		log.Fatal(err)
	}
}
