package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mjcarpent/calc-app/handler"
	"github.com/mjcarpent/calc-lib"
)

func main() {

	var op string
	flag.StringVar(&op, "op", "+", "Calculator operation to perform")
	flag.Parse()

	operator, ok := calculators[op]
	if !ok {
		fmt.Println("Invalid operator specified[" + op + "]")
		os.Exit(1)
	}

	calcHandler := handler.NewCLIHandler(os.Stdout, operator)
	err := calcHandler.Handle(flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

var calculators = map[string]handler.Calculator{

	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
	"%": &calc.Modulo{},
}
