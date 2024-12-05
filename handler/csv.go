package handler

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

type CSVHandler struct {
	reader *csv.Reader
	writer *csv.Writer
	logger *log.Logger
}

func NewCSVHandler(reader io.Reader, writer io.Writer, logger *log.Logger) *CSVHandler {
	return &CSVHandler{reader: csv.NewReader(reader), writer: csv.NewWriter(writer), logger: logger}
}

func (this *CSVHandler) Handle() error {

	for {

		record, err := this.reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			this.logger.Printf("%s: [%s]\n", csvReadError, err)
			return csvReadError
		}

		operand1, err := strconv.Atoi(record[0])
		if err != nil {
			this.logger.Printf("%s[%s]\n", operandMustBeInt, record[0])
			continue
		}

		operation, ok := Calculators[record[1]]
		if !ok {
			this.logger.Printf("%s[%s]\n", unknownOperation, record[1])
			continue
		}

		operand2, err := strconv.Atoi(record[2])
		if err != nil {
			this.logger.Printf("%s[%s]\n", operandMustBeInt, record[2])
			continue
		}

		answer := operation.Calculate(operand1, operand2)
		err = this.writer.Write(append(record, fmt.Sprint(answer)))
		if err != nil {
			this.logger.Printf("%v: %v", csvWriteError, err)
			return fmt.Errorf("%w: %w", csvWriteError, err)
		}
	}

	this.writer.Flush()
	return this.writer.Error()
}
