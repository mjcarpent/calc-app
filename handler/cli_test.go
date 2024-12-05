package handler

import (
	"errors"
	"testing"

	"github.com/mjcarpent/calc-lib"
)

func TestCLIHandler_BadFirstParam(t *testing.T) {

	cliHandler := NewCLIHandler(calc.Addition{})
	_, err := cliHandler.Handle([]string{"a", "2"})
	if err == nil {
		t.Errorf("Expected error and got success")
		return
	}

	if !errors.Is(err, operandMustBeInt) {
		t.Errorf("The returned error[%v] was not the expected error[%v]", err, operandMustBeInt)
	}
}

func TestCLIHandler_BadSecondParam(t *testing.T) {

	cliHandler := NewCLIHandler(calc.Addition{})
	_, err := cliHandler.Handle([]string{"5", "t"})
	if err == nil {
		t.Errorf("Expected error and got success")
		return
	}

	if !errors.Is(err, operandMustBeInt) {
		t.Errorf("The returned error[%v] was not the expected error[%v]", err, operandMustBeInt)
	}
}

func TestCLIHandler_GoodOperator(t *testing.T) {

	tests := []struct {
		name       string
		calculator Calculator
		args       []string
		want       int
		wantErr    bool
	}{
		{name: "Good addition operation", calculator: calc.Addition{}, args: []string{"5", "3"}, want: 8, wantErr: false},
		{name: "Good subtraction operation", calculator: calc.Subtraction{}, args: []string{"3", "8"}, want: -5, wantErr: false},
		{name: "Good multiplication operation", calculator: calc.Multiplication{}, args: []string{"5", "5"}, want: 25, wantErr: false},
		{name: "Good division operation", calculator: calc.Division{}, args: []string{"17", "3"}, want: 5, wantErr: false},
		{name: "Good modulus operation", calculator: calc.Modulus{}, args: []string{"17", "3"}, want: 2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cliHandler := NewCLIHandler(tt.calculator)
			got, err := cliHandler.Handle(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
