package calculator

type Calculator interface {
	SetLeftOperand(o string) error
	SetRightOperand(o string) error
	SetOperator(o string) error
	Calculate() error
	PrintResult() string
}
