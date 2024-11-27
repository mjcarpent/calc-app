package main

import (
	"log"
	"os"

	"github.com/mjcarpent/calc-app/handler"
	"github.com/mjcarpent/calc-lib"
)

func main() {

	handler := handler.NewCSVHandler(os.Stdin, os.Stdout, os.Stderr, calculators)
	err := handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]handler.Calculator{

	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
	"%": &calc.Modulo{},
}
