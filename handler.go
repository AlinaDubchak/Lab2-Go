package lab2

import (
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
  inputBytes, err := io.ReadAll(ch.Input)
  if err != nil {
    return err
  }

  expr := string(inputBytes)
  trimmedExpr := strings.TrimSpace(expr)
  result, err := SolvePostfixExpression(trimmedExpr)
  if err != nil {
    return fmt.Errorf("error solving expression: %v", err)
  }

  _, err = fmt.Fprintf(ch.Output, "%s\n", []byte(result))
  if err != nil {
    return err
  }

  return nil
}