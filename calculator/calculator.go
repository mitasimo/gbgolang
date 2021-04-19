package calculator

import (
	"errors"
)

var (
	ErrSecondOperandIsZero = errors.New("второй операнд не может быть равен ноль (0)")
	ErrUnknownOperation    = errors.New("неизвестная операция")
)

type Operation struct {
	First, Second int
	Operator      string
}

func (o *Operation) Calculate() (r int, err error) {

	switch o.Operator {
	case "+":
		r = o.add()
	case "-":
		r = o.subtract()
	case "*":
		r = o.multiply()
	case "/":
		if o.Second == 0 {
			err = ErrSecondOperandIsZero
		} else {
			r = o.divide()
		}
	case "%":
		if o.Second == 0 {
			err = ErrSecondOperandIsZero
		} else {
			r = o.rd()
		}
	default:
		err = ErrUnknownOperation
	}

	return
}

func (op *Operation) add() int {
	return op.First + op.Second
}

func (op *Operation) subtract() int {
	return op.First - op.Second
}

func (op *Operation) multiply() int {
	return op.First * op.Second
}

func (op *Operation) divide() int {
	return op.First / op.Second
}

// остаток от деления
func (op *Operation) rd() int {
	return op.First % op.Second
}
