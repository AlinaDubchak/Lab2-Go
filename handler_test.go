package lab2

import (
	"bytes"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestComputeHandler(c *C) {
	input := "5 2 -"
	output := &bytes.Buffer{}

	handler := ComputeHandler{
		Input:  strings.NewReader(input),
		Output: output,
	}

	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(output.String(), Equals, "3.00\n")
}

func (s *MySuite) TestComputeHandlerError(c *C) {
	input := ""
	output := &bytes.Buffer{}

	handler := ComputeHandler{
		Input:  strings.NewReader(input),
		Output: output,
	}

	err := handler.Compute()
	c.Assert(err, ErrorMatches, "EOF")
}