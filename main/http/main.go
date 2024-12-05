package main

import (
	"log"
	"net/http"

	"github.com/mjcarpent/calc-app/handler"
)

func main() {

	log.Println("Listening ...")
	err := http.ListenAndServe("localhost:8080", handler.NewHTTPRouter())
	if err != nil {
		log.Fatalln(err)
	}
}
