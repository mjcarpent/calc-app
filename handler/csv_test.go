package handler

import (
	"bytes"
	"strings"
	"testing"
)

func assertStdOut(t *testing.T, received, expected string) {

	t.Helper()
	if received != expected {
		t.Errorf("The expected output[%s] does not match the received output[%s]", expected, received)
	}
}

func assertStdErr(t *testing.T, received, expected string) {

	t.Helper()
	received = strings.Trim(received, "\n")
	if received != expected {
		t.Errorf("The expected output[%s] does not match the received output[%s]", expected, received)
	}
}

func assertEmpty(t *testing.T, received string) {

	t.Helper()
	received = strings.Trim(received, "\n")
	if received != "" {
		t.Errorf("The output buffer is not empty[%s]", received)
	}
}

func assertOutputContainsString(t *testing.T, received, expected string) {

	t.Helper()
	if !strings.Contains(received, expected) {
		t.Errorf("The expected output[%s] is not found in the received output[%s]", expected, received)
	}
}

func TestNewCSVHandler_SingleLineInput(t *testing.T) {

	stdin := "1,+,3\n"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	NewCSVHandler(strings.NewReader(stdin), &stdout, &stderr, calculators).Handle()
	assertStdOut(t, strings.Trim(stdout.String(), "\n"), strings.Trim(stdin, "\n")+",4")
	assertEmpty(t, stderr.String())
}

func TestNewCSVHandler_MultiLineInput(t *testing.T) {

	stdin := "1,+,3\n5,-,-3\n4,%,4\n4,/,4\n5,*,5"
	expectedOutput := "1,+,3,4\n5,-,-3,8\n4,%,4,0\n4,/,4,1\n5,*,5,25"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	NewCSVHandler(strings.NewReader(stdin), &stdout, &stderr, calculators).Handle()
	assertStdOut(t, strings.Trim(stdout.String(), "\n"), expectedOutput)
	assertEmpty(t, stderr.String())
}

func TestNewCSVHandler_SingleLineInputBad(t *testing.T) {

	stdin := "1,2,3"
	expectedOutput := "Line[0] parameter 2[2] is not a valid operator"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	NewCSVHandler(strings.NewReader(stdin), &stdout, &stderr, calculators).Handle()
	assertOutputContainsString(t, strings.Trim(stderr.String(), "\n"), expectedOutput)
	assertEmpty(t, stdout.String())
}

func TestNewCSVHandler_MultiLineInputBad(t *testing.T) {

	stdin := "1,+,3\nNaN,-,-3\n4,%,#\n4,/,4\n5,*,5"
	expectedOutput := "1,+,3,4\n4,/,4,1\n5,*,5,25"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	NewCSVHandler(strings.NewReader(stdin), &stdout, &stderr, calculators).Handle()
	assertStdOut(t, strings.Trim(stdout.String(), "\n"), expectedOutput)
	assertOutputContainsString(t, strings.Trim(stderr.String(), "\n"), "Line[1] parameter 1 failed with the following message[strconv.Atoi: parsing \"NaN\": invalid syntax]")
	assertOutputContainsString(t, strings.Trim(stderr.String(), "\n"), "Line[2] parameter 3 failed with the following message[strconv.Atoi: parsing \"#\": invalid syntax]")
}

func TestNewCSVHandler_BadReader(t *testing.T) {

	stdin := &ReadError{}
	//expectedOutput := "1,+,3,4\n4,/,4,1\n5,*,5,25"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	NewCSVHandler(stdin, &stdout, &stderr, calculators).Handle()
	//assertStdOut(t, strings.Trim(stdout.String(), "\n"), expectedOutput)
	assertOutputContainsString(t, strings.Trim(stderr.String(), "\n"), "Error reading from stdin")
}

type ReadError struct {
	err error
}

func (this *ReadError) Read(p []byte) (n int, err error) {
	return 0, this.err
}
