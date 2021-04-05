package main

import (
	"fmt"
)

func main() {
	var (
		first, second, result int
		operator              string
		err                   error
	)

	for {
		fmt.Print("Введите значение первого операнда: ")
		_, err = fmt.Scanln(&first)
		if err != nil {
			fmt.Printf("Ошибка ввода первого операнда: %v\n", err)
			continue
		}
		fmt.Print("Введите значение второго операнда: ")
		_, err = fmt.Scanln(&second)
		if err != nil {
			fmt.Printf("Ошибка ввода второго операнда: %v\n", err)
			continue
		}
		fmt.Print("Введите оператор (+, -, *, /, %): ")
		_, err = fmt.Scanln(&operator)
		if err != nil {
			fmt.Printf("Ошибка ввода оператора: %v\n", err)
			continue
		}

		switch operator {
		case "+":
			result = first + second
		case "-":
			result = first - second
		case "*":
			result = first * second
		case "/":
			if second == 0.0 {
				fmt.Println("Ошибка: второй оперант не может быть равен ноль (0)")
				continue
			}
			result = first / second
		case "%":
			if second == 0.0 {
				fmt.Println("Ошибка: второй оперант не может быть равен ноль (0)")
				continue
			}
			result = first % second
		default:
			fmt.Printf("Ошибка: оператор %s не поддерживается\n", operator)
			continue
		}
		fmt.Printf("Результат %d %s %d = %d\n\n", first, operator, second, result)
	}
}
