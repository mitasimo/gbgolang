package intposcalc

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mitasimo/gbgolang/calculator/common"
)

var ErrNonPositive = errors.New("отрицательное число")

type Calculator struct {
	leftOperand, rightOperand, result int
	operator                          string
	err                               error
}

func (c *Calculator) SetLeftOperand(o string) error {
	v, err := strconv.ParseInt(o, 10, 64)
	if err != nil {
		return err
	}
	if v < 0 {
		return ErrNonPositive
	}
	c.leftOperand = int(v)
	return nil
}

func (c *Calculator) SetRightOperand(o string) error {
	v, err := strconv.ParseInt(o, 10, 64)
	if err != nil {
		return err
	}
	if v < 0 {
		return ErrNonPositive
	}
	c.rightOperand = int(v)
	return nil
}

func (c *Calculator) SetOperator(o string) error {
	return common.SetOperator(o, &c.operator)
}

func (c *Calculator) Calculate() error {
	switch c.operator {
	case "+":
		c.result = c.leftOperand + c.rightOperand
	case "-":
		c.result = c.leftOperand - c.rightOperand
	case "*":
		c.result = c.leftOperand * c.rightOperand
	case "/":
		if c.rightOperand == 0 {
			c.err = errors.New("правый операнд равен 0(ноль)")
			return c.err
		}
		c.result = c.leftOperand / c.rightOperand
	case "%":
		if c.rightOperand == 0 {
			c.err = errors.New("правый операнд равен 0(ноль)")
			return c.err
		}
		c.result = c.leftOperand % c.rightOperand
	}
	return nil
}

func (c *Calculator) PrintResult() string {
	if c.err != nil {
		return fmt.Sprintf("ошибка: %v", c.err)
	}
	return fmt.Sprintf("%d %s %d = %d", c.leftOperand, c.operator, c.rightOperand, c.result)
}
