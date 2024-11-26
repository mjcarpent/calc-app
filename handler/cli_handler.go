package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface {
	Calculate(a, b int) int
}

type CLIHandler struct {
	writer     io.Writer
	calculator Calculator
}

func NewCLIHandler(stdout io.Writer, calculator Calculator) *CLIHandler {
	return &CLIHandler{
		writer:     stdout,
		calculator: calculator,
	}
}

func (this *CLIHandler) Handle(args []string) error {

	if len(args) != 2 {
		return insufficientArgs
	}

	param1, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", invalidArg, err)
	}

	param2, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", invalidArg, err)
	}

	param3 := this.calculator.Calculate(param1, param2)
	_, err = fmt.Fprintf(this.writer, "%d", param3)
	if err != nil {
		return err
	}

	return nil
}

var (
	insufficientArgs = errors.New("Usage: Addition requires 2 operands")
	invalidArg       = errors.New("invalid argument")
)
