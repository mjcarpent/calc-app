package handler

import (
	"bytes"
	"errors"
	"testing"

	"github.com/mjcarpent/calc-lib"
)

func assertError(t *testing.T, err, expected error) {

	t.Helper()
	if !errors.Is(err, expected) {
		t.Errorf("expected err to wrap %v, but it didn't", expected)
	}
}

func assertOutput(t *testing.T, output string, expected string) {

	t.Helper()
	if output != expected {
		t.Errorf("The expected output[%s] does not match the output[%s] received", output, expected)
	}
}

func TestHandler_NoParameters(t *testing.T) {

	handler := NewCLIHandler(nil, nil)
	err := handler.Handle(nil)
	assertError(t, err, insufficientArgs)
}

func TestHandler_OperandOneBad(t *testing.T) {

	handler := NewCLIHandler(new(bytes.Buffer), &calc.Addition{})
	err := handler.Handle([]string{"a", "2"})
	assertError(t, err, invalidArg)
}

func TestHandler_OperandTwoBad(t *testing.T) {

	handler := NewCLIHandler(new(bytes.Buffer), &calc.Addition{})
	err := handler.Handle([]string{"7", "z"})
	assertError(t, err, invalidArg)
}

func TestHandler_Success(t *testing.T) {

	buf := new(bytes.Buffer)
	handler := NewCLIHandler(buf, &calc.Addition{})
	_ = handler.Handle([]string{"7", "3"})
	assertOutput(t, buf.String(), "10")
}

func TestHandler_BadOutput(t *testing.T) {

	myError := errors.New("splat")
	output := &WriterError{err: myError}
	handler := NewCLIHandler(output, &calc.Addition{})
	err := handler.Handle([]string{"7", "5"})
	assertError(t, err, myError)
}

type WriterError struct {
	err error
}

func (this *WriterError) Write(p []byte) (n int, err error) {
	return 0, this.err
}
