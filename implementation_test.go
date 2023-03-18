package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func TestEvaluate(t *testing.T) {
	TestingT(t)
}

func (s *MySuite) TestSimple(c *C) {
	expressions := []struct {
		expr   string
		wanted string
	}{
		{"2 3 +", "5.00"},
		{"5 2 -", "3.00"},
		{"4 3 *", "12.00"},
		{"6 2 /", "3.00"},
	}
	for _, tt := range expressions {
		got, err := SolvePostfixExpression(tt.expr)
		c.Assert(err, IsNil)
		c.Assert(got, Equals, tt.wanted)
	}
}

func (s *MySuite) TestComplex(c *C) {
	expressions := []struct {
		expr   string
		wanted string
	}{
		{"1 2 + 3 * 4 /", "2.25"},
		{"2 3 + 4 * 5 -", "15.00"},
		{"2 3 ^ 4 + 5 - 6 /", "1.17"},
	}
	for _, tt := range expressions {
		got, err := SolvePostfixExpression(tt.expr)
		c.Assert(err, IsNil)
		c.Assert(got, Equals, tt.wanted)
	}
}

func (s *MySuite) TestInvalid(c *C) {
	expressions := []struct {
		expr string
	}{
		{""},
		{"2 + 3"},
		{"2 x 3 +"},
	}
	for _, tt := range expressions {
		_, err := SolvePostfixExpression(tt.expr)
		c.Assert(err, NotNil)
		c.Assert(err, ErrorMatches, "invalid expression")
	}
}

func ExampleSolvePostfixExpression() {
	result, err := SolvePostfixExpression("2 5 ^ 6 3 * + 5 - 5 /")
	if err == nil {
		fmt.Println(result)
	} else {
		panic(err)
	}
	// Output: 9.00
}
