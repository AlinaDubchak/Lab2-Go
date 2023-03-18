package lab2

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func NewComputeHandler(input io.Reader, output io.Writer) ComputeHandler {
	return ComputeHandler{input, output}
}

func (ch *ComputeHandler) Compute() error {
	buffer := make([]byte, 128)
	_, err := ch.Input.Read(buffer)
	if err != nil {
		return err
	}
	buffer = bytes.Trim(buffer, "\x00")

	expr := string(buffer)
	trimeExpr := strings.Trim(expr, " \n")
	result, err := SolvePostfixExpression(trimeExpr)
	if err != nil {
		return fmt.Errorf("error solving expression: %v", err)
	}
	_, err = fmt.Fprintf(ch.Output, "%s", []byte(result))
	if err != nil {
		return err
	}
	return nil
}
