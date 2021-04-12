package main

import (
	"github.com/mitasimo/gbgolang/calculator"
)

func main() {

	actions := []action{
		{"Калькулятор", calculator.RunCalculator},
		{"Фибоначчи", RunFibonachi},
	}
	menu(actions)

}
