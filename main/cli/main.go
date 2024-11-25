package main

import (
	"fmt"
	"os"

	"github.com/mjcarpent/calc-app/handler"
	"github.com/mjcarpent/calc-lib"
)

func main() {

	add := handler.NewCLIHandler(os.Stdout, &calc.Addition{})
	err := add.Handle(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
