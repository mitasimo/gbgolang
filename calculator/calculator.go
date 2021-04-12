package calculator

import "fmt"

func RunCalculator() {
	const (
		ErrSecondOperandIsZero = "Ошибка: второй операнд не может быть равен ноль (0)"
	)

	type Operation struct {
		first, second, result int
		operator              string
	}

	var (
		op  Operation
		err error
	)

	fmt.Println()
	fmt.Print("Введите значение первого операнда: ")
	_, err = fmt.Scanln(&op.first)
	if err != nil {
		fmt.Printf("\nОшибка ввода первого операнда: %v\n", err)
		return
	}
	fmt.Print("Введите значение второго операнда: ")
	_, err = fmt.Scanln(&op.second)
	if err != nil {
		fmt.Printf("\nОшибка ввода второго операнда: %v\n", err)
		return
	}
	fmt.Print("Введите оператор (+, -, *, /, %): ")
	_, err = fmt.Scanln(&op.operator)
	if err != nil {
		fmt.Printf("\nОшибка ввода оператора: %v\n", err)
		return
	}

	switch op.operator {
	case "+":
		op.result = op.first + op.second
	case "-":
		op.result = op.first - op.second
	case "*":
		op.result = op.first * op.second
	case "/":
		if op.second == 0.0 {
			fmt.Println(ErrSecondOperandIsZero)
			return
		}
		op.result = op.first / op.second
	case "%":
		if op.second == 0.0 {
			fmt.Println(ErrSecondOperandIsZero)
			return
		}
		op.result = op.first % op.second
	default:
		fmt.Printf("Ошибка: оператор %s не поддерживается\n", op.operator)
		return
	}
	fmt.Printf("Результат %d %s %d = %d\n", op.first, op.operator, op.second, op.result)
}
