package main

import (
	"fmt"
	"os"

	"github.com/mitasimo/gbgolang/calculator"
	"github.com/mitasimo/gbgolang/calculator/floatcalc"
	"github.com/mitasimo/gbgolang/calculator/intcalc"
	"github.com/mitasimo/gbgolang/calculator/intposcalc"
)

func main() {

	var (
		calc     calculator.Calculator
		calcType int
	)

	fmt.Println("0 - калькулятор целых числел")
	fmt.Println("1 - калькулятор целых положительных чисел")
	fmt.Println("2 - калькулятор вещественных чисел")
	fmt.Print("Какой калькулятор использовать: ")
	fmt.Scanln(&calcType)

	switch calcType {
	case 0:
		calc = new(intcalc.Calculator)
	case 1:
		calc = new(intposcalc.Calculator)
	case 2:
		calc = new(floatcalc.Calculator)
	default:
		fatal("неизвестный тип калькулятора")
	}

	var input string

	fmt.Print("Введите первый операнд: ")
	fmt.Scanln(&input)
	if err := calc.SetLeftOperand(input); err != nil {
		fatal(err.Error())
	}
	fmt.Print("Введите второй операнд: ")
	fmt.Scanln(&input)
	if err := calc.SetRightOperand(input); err != nil {
		fatal(err.Error())
	}
	fmt.Print("Введите оператор (+, -, *, /, %): ")
	fmt.Scanln(&input)
	if err := calc.SetOperator(input); err != nil {
		fatal(err.Error())
	}
	if err := calc.Calculate(); err != nil {
		fatal(err.Error())
	}

	fmt.Println(calc.PrintResult())
}

// Выводит сообщение в поток ошибок и завершает выполнение
func fatal(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
