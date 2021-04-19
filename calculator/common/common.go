package common

import "errors"

var ErrUnknownOperator = errors.New("неизвестный оператор")

func SetOperator(o string, s *string) error {
	switch o {
	case "+", "-", "*", "/", "%":
		*s = o
	default:
		return ErrUnknownOperator
	}
	return nil
}
