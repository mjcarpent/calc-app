package handler

import (
	"strconv"
)

type CLIHandler struct {
	calculator Calculator
}

func NewCLIHandler(calculator Calculator) *CLIHandler {
	return &CLIHandler{calculator: calculator}
}

func (this *CLIHandler) Handle(args []string) (int, error) {

	param1, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, operandMustBeInt
	}

	param2, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, operandMustBeInt
	}

	return this.calculator.Calculate(param1, param2), nil
}
