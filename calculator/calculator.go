package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Calculator struct {
	Number float64
	Ops    [][]string
}

func NewCalculator(initval float64) *Calculator {
	return &Calculator{
		Number: initval,
		Ops:    [][]string{},
	}
}

func (c *Calculator) Exec(input []string) (float64, error) {
	if len(input) < 1 {
		return c.Number, fmt.Errorf("Invalid input: empty input")
	}

	num, err := c.exec(input)
	if err != nil {
		return num, err
	}

	if input[0] != "repeat" {
		c.Ops = append(c.Ops, input)
	}

	return num, nil
}

func (c *Calculator) exec(input []string) (float64, error) {
	var (
		n   float64
		err error
	)

	if strings.ToLower(input[0]) == "add" ||
		strings.ToLower(input[0]) == "subtract" ||
		strings.ToLower(input[0]) == "multiply" ||
		strings.ToLower(input[0]) == "divide" ||
		strings.ToLower(input[0]) == "repeat" {
		if len(input) < 2 {
			return c.Number, fmt.Errorf("Invalid input: Operation `%s` requires 1 argument.", input[0])
		}

		n, err = strconv.ParseFloat(input[1], 10)
		if err != nil {
			return c.Number, fmt.Errorf("Invalid input: Operation `%s` requires numeric argument.", input[0])
		}
	}

	switch strings.ToLower(input[0]) {
	case "add":
		c.add(n)
	case "subtract":
		c.subtract(n)
	case "multiply":
		c.multiply(n)
	case "divide":
		err = c.divide(n)
	case "cancel":
		c.cancel()
	case "abs":
		c.abs()
	case "neg":
		c.neg()
	case "sqrt":
		c.sqrt()
	case "sqr":
		c.sqr()
	case "cubert":
		c.cubert()
	case "cube":
		c.cube()
	case "repeat":
		err = c.repeat(n)
	default:
		return c.Number, fmt.Errorf("Unknown operation: %s", input[0])
	}
	if err != nil {
		return c.Number, err
	}

	return c.Number, nil
}

func (c *Calculator) add(n float64) {
	c.Number += float64(n)
}

func (c *Calculator) subtract(n float64) {
	c.Number -= float64(n)
}

func (c *Calculator) multiply(n float64) {
	c.Number *= float64(n)
}

func (c *Calculator) divide(n float64) error {
	if n == 0 {
		return fmt.Errorf("Invalid input: Operation `Divide` does not accept 0 as argument")
	}
	c.Number /= float64(n)

	return nil
}

func (c *Calculator) cancel() {
	c.Number = 0
	c.Ops = [][]string{}
}

func (c *Calculator) abs() {
	if c.Number < 0 {
		c.neg()
	}
}

func (c *Calculator) neg() {
	c.Number *= -1
}

func (c *Calculator) sqrt() {
	c.Number = math.Sqrt(c.Number)
}

func (c *Calculator) sqr() {
	c.Number *= c.Number
}

func (c *Calculator) cubert() {
	c.Number = math.Cbrt(c.Number)
}

func (c *Calculator) cube() {
	c.Number = c.Number * c.Number * c.Number
}

func (c *Calculator) repeat(n float64) error {
	if n < 1 {
		return fmt.Errorf("Invalid input: Operation `repeat` must have positive numeric argument, got %d", int(n))
	}

	if len(c.Ops) < int(n) {
		return fmt.Errorf("Invalid input: Operation `repeat` cannot be executed. Current number of operation is %d, got %d", len(c.Ops), int(n))
	}

	repeatedOps := [][]string{}
	for _, v := range c.Ops {
		_, err := c.exec(v)
		if err != nil {
			return fmt.Errorf("Error while executing operation: %w", err)
		}
		repeatedOps = append(repeatedOps, v)
	}

	c.Ops = append(c.Ops, repeatedOps...)

	return nil
}
