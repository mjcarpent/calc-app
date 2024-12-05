package handler

import (
	"errors"

	"github.com/mjcarpent/calc-lib"
)

type Calculator interface {
	Calculate(a, b int) int
}

var Calculators = map[string]Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
	"%": &calc.Modulus{},
}

var operators = map[string]string{
	"add": "+",
	"sub": "-",
	"mul": "*",
	"div": "/",
	"mod": "%",
}

var (
	unknownOperation = errors.New("Unknown operation")
	operandMustBeInt = errors.New("The given operand must be an integer")
	csvReadError     = errors.New("Unknown failure returned on read to csv")
	csvWriteError    = errors.New("Unknown failure returned on write to csv")
)
