package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mjcarpent/calc-app/handler"
)

func main() {

	csvHandler := handler.NewCSVHandler(os.Stdin, os.Stdout, log.New(os.Stderr, "", 0))
	err := csvHandler.Handle()
	if err != nil {
		fmt.Printf("Failure performing operation: %v\n", err)
		os.Exit(1)
	}
}
