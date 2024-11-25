package main

import (
	"fmt"
	"os"

	"github.com/mcarpenter/calc-app/handler"
)

func main() {

	var a = handler.CLIHandler{os.Stdout, calc.Addition{}}
	err := a.Handle(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
