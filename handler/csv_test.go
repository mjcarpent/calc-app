package handler

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestCSVFixture(t *testing.T) {
	gunit.Run(new(CSVTestFixture), t)
}

type CSVTestFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.

	// Declare useful state here (probably the stuff being tested, any fakes, etc...).
}

func (this *CSVTestFixture) SetupStuff() {
}

func (this *CSVTestFixture) TeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
}

func (this *CSVTestFixture) Test_GoodSingleLineAdd() {

	reader := bytes.NewReader([]byte("2,+,3\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	this.So(output, should.Equal, "2,+,3,5")
}

func (this *CSVTestFixture) Test_GoodSingleLineSubtract() {

	reader := bytes.NewReader([]byte("2,-,3\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	this.So(output, should.Equal, "2,-,3,-1")
}

func (this *CSVTestFixture) Test_GoodSingleLineMultiply() {

	reader := bytes.NewReader([]byte("3,*,3\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	this.So(output, should.Equal, "3,*,3,9")
}

func (this *CSVTestFixture) Test_GoodSingleLineDivide() {

	reader := bytes.NewReader([]byte("17,/,3\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	this.So(output, should.Equal, "17,/,3,5")
}

func (this *CSVTestFixture) Test_GoodSingleLineModulus() {

	reader := bytes.NewReader([]byte("17,%,3\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	this.So(output, should.Equal, "17,%,3,2")
}

func (this *CSVTestFixture) Test_BadFirstOperand() {

	reader := bytes.NewReader([]byte("a,+,3\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	//output := strings.TrimSpace(writer.String())
	if !strings.Contains(logOutput.String(), fmt.Sprintf("%v", operandMustBeInt)) {
		fmt.Printf("Output[%s] does not contain expected[%s]", logOutput.String(), fmt.Sprintf("%v", operandMustBeInt))
	}
}

func (this *CSVTestFixture) Test_BadSecondOperand() {

	reader := bytes.NewReader([]byte("2,+,b\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	//output := strings.TrimSpace(writer.String())
	if !strings.Contains(logOutput.String(), fmt.Sprintf("%v", operandMustBeInt)) {
		fmt.Printf("Output[%s] does not contain expected[%s]", logOutput.String(), fmt.Sprintf("%v", operandMustBeInt))
	}
}

func (this *CSVTestFixture) Test_BadOperator() {

	reader := bytes.NewReader([]byte("2,@,3\n"))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	//output := strings.TrimSpace(writer.String())
	if !strings.Contains(logOutput.String(), fmt.Sprintf("%v", unknownOperation)) {
		fmt.Printf("Output[%s] does not contain expected[%s]", logOutput.String(), fmt.Sprintf("%v", unknownOperation))
	}
}

func (this *CSVTestFixture) Test_GoodMultilineInput() {

	data := "1,+,2\n3,-,4\n5,*,6\n7,/,8\n9,%,2\n"
	reply := "1,+,2,3\n3,-,4,-1\n5,*,6,30\n7,/,8,0\n9,%,2,1"
	reader := bytes.NewReader([]byte(data))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	this.So(output, should.Equal, reply)
}

func (this *CSVTestFixture) Test_AllLinesBad() {

	data := "a,+,2\n3,@,4\n5,*,y\n/,7,8\n9,2,%\n"
	reader := bytes.NewReader([]byte(data))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	if len(output) != 0 {
		fmt.Printf("Output[%s] given when none was expected", output)
	}

	loggerString := logOutput.String()
	mustBeInt := fmt.Sprintf("%s", operandMustBeInt)
	unknownOp := fmt.Sprintf("%s", unknownOperation)

	if !strings.Contains(loggerString, mustBeInt+"[a]") {
		fmt.Printf("Logger is missing error for int conversion[%s]", loggerString)
	}

	if !strings.Contains(loggerString, unknownOp+"[@]") {
		fmt.Printf("Logger is missing error for the bad operator[%s]", loggerString)
	}

	if !strings.Contains(loggerString, mustBeInt+"[y]") {
		fmt.Printf("Logger is missing error for int conversion[%s]", loggerString)
	}

	if !strings.Contains(loggerString, mustBeInt+"[/]") {
		fmt.Printf("Logger is missing error for int conversion[%s]", loggerString)
	}
	if !strings.Contains(loggerString, unknownOp+"[2]") {
		fmt.Printf("Logger is missing error for the bad operator[%s]", loggerString)
	}
}

func (this *CSVTestFixture) Test_FirstLineBad() {

	data := "3,+,t\n3,*,4\n5,*,2\n8,/,1"
	response := "3,*,4,12\n5,*,2,10\n8,/,1,8"
	reader := bytes.NewReader([]byte(data))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	if output != response {
		fmt.Printf("Output[%s] given did not match the expected[%s]", output, response)
	}

	loggerString := logOutput.String()
	mustBeInt := fmt.Sprintf("%s", operandMustBeInt)

	if !strings.Contains(loggerString, mustBeInt+"[t]") {
		fmt.Printf("Logger is missing error for int conversion[%s]", loggerString)
	}
}

func (this *CSVTestFixture) Test_LastLineBad() {

	data := "3,+,14\n3,*,4\n5,*,2\nz,/,1"
	response := "3,+,14,17\n3,*,4,12\n5,*,2,10"
	reader := bytes.NewReader([]byte(data))
	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	if output != response {
		fmt.Printf("Output[%s] given did not match the expected[%s]", output, response)
	}

	loggerString := logOutput.String()
	mustBeInt := fmt.Sprintf("%s", operandMustBeInt)

	if !strings.Contains(loggerString, mustBeInt+"[z]") {
		fmt.Printf("Logger is missing error for int conversion[%s]", loggerString)
	}
}

func (this *CSVTestFixture) Test_MiddleLineBad() {

	data := "3,+,14\nq,*,4\n6,/,1"
	response := "3,+,14,17\n6,/,1,6"
	reader := bytes.NewReader([]byte(data))

	writer := &bytes.Buffer{}

	var logOutput bytes.Buffer
	logger := log.New(&logOutput, "", 0)

	handler := NewCSVHandler(reader, writer, logger)
	err := handler.Handle()
	if err != nil {
		fmt.Printf("Unexpected error returned from test[%v]", err)
	}

	output := strings.TrimSpace(writer.String())
	if output != response {
		fmt.Printf("Output[%s] given did not match the expected[%s]", output, response)
	}

	loggerString := logOutput.String()
	mustBeInt := fmt.Sprintf("%s", operandMustBeInt)

	if !strings.Contains(loggerString, mustBeInt+"[q]") {
		fmt.Printf("Logger is missing error for int conversion[%s]", loggerString)
	}
}
