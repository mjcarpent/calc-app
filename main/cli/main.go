package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mjcarpent/calc-app/handler"
)

func main() {

	var op string
	flag.StringVar(&op, "op", "+", "The math operation to perform")
	flag.Parse()

	operator, ok := handler.Calculators[op]
	if !ok {
		fmt.Printf("Unknown operation specified[%s]\n", op)
		os.Exit(1)
	}

	handle := handler.NewCLIHandler(operator)
	answer, err := handle.Handle(flag.Args())
	if err != nil {
		fmt.Printf("Failure performing operation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", answer)
}
