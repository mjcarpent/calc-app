package handler

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/mjcarpent/calc-lib"
)

type CSVHandler struct {
	stdin       *csv.Reader
	stdout      *csv.Writer
	stderr      *log.Logger
	calculators map[string]Calculator
}

func NewCSVHandler(stdin io.Reader, stdout, stderr io.Writer, calculators map[string]Calculator) *CSVHandler {
	return &CSVHandler{
		stdin:       csv.NewReader(stdin),
		stdout:      csv.NewWriter(stdout),
		stderr:      log.New(stderr, "", log.LstdFlags),
		calculators: calculators,
	}
}

func (this *CSVHandler) Handle() error {

	for i := 0; ; i++ {

		record, err := this.stdin.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			this.stderr.Printf("Error reading from stdin[%s]", err)
			return err
		}

		if len(record) < 0 {
			this.stderr.Printf("Line[%d] is an empty line\n", i)
			continue
		}

		param1, err := strconv.Atoi(record[0])
		if err != nil {
			this.stderr.Printf("Line[%d] parameter 1 failed with the following message[%s]\n", i, err)
			continue
		}

		op := record[1]
		_, ok := calculators[op]
		if !ok {
			this.stderr.Printf("Line[%d] parameter 2[%s] is not a valid operator\n", i, op)
			continue
		}

		calculator := this.calculators[op]
		param2, err := strconv.Atoi(record[2])
		if err != nil {
			this.stderr.Printf("Line[%d] parameter 3 failed with the following message[%s]\n", i, err)
			continue
		}

		answer := calculator.Calculate(param1, param2)
		err = this.stdout.Write(append(record, strconv.Itoa(answer)))
		if err != nil {
			this.stderr.Printf("Line[%d] unable to write result to output[%s]\n", i, err)
		} else {
			this.stdout.Flush()
		}
	}

	return nil
}

var calculators = map[string]Calculator{

	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
	"%": &calc.Modulo{},
}
