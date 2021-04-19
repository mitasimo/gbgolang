package main

import (
	"fmt"
	"log"

	calc "github.com/mitasimo/gbgolang/calculator"
)

func main() {
	RunCalculator()
}

func RunCalculator() {

	var (
		op  calc.Operation
		err error
	)

	fmt.Println()
	fmt.Print("Введите значение первого операнда: ")
	_, err = fmt.Scanln(&op.First)
	if err != nil {
		fmt.Printf("\nОшибка ввода первого операнда: %v\n", err)
		return
	}
	fmt.Print("Введите значение второго операнда: ")
	_, err = fmt.Scanln(&op.Second)
	if err != nil {
		fmt.Printf("\nОшибка ввода второго операнда: %v\n", err)
		return
	}
	fmt.Print("Введите оператор (+, -, *, /, %): ")
	_, err = fmt.Scanln(&op.Operator)
	if err != nil {
		fmt.Printf("\nОшибка ввода оператора: %v\n", err)
		return
	}

	r, err := op.Calculate()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Результат %d %s %d = %d\n", op.First, op.Operator, op.Second, r)
}
