package main

import (
	"fmt"
	"math"
)

type action struct {
	name string
	do   func()
}

func menu(actions []action) {
	var choice int

	for {
		choice = 0

		fmt.Println("Действия:")
		for c, a := range actions {
			fmt.Printf("%d: %s\n", c+1, a.name)
		}
		fmt.Print("Введите номер действия (0 - выход): ")
		n, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}
		if n == 0 {
			fmt.Println("Ошибка: не выбрано действие")
			return
		}

		if choice == 0 {
			break
		}
		if choice > len(actions) {
			fmt.Printf("Ошибка: действие с номером %d отсутствует\n", choice)
		} else {
			actions[choice-1].do()
		}
	}
}

func main() {
	actions := []action{
		{"Площадь прямоугольника", square},
		{"Диаметр круга и длина окружности", circle},
		{"Разряды числа", digits},
	}
	menu(actions)
}

func square() {
	var x, y int
	fmt.Println("Действие: расчет площади прямоугольника")

	fmt.Print("Введите ширину: ")
	n, err := fmt.Scan(&x)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	if n == 0 {
		fmt.Println("Ошибка: не введено значение ширины")
		return
	}
	fmt.Print("Введите длину: ")
	n, err = fmt.Scan(&y)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	if n == 0 {
		fmt.Println("Ошибка: не введено значение длины")
		return
	}

	fmt.Printf("Площать прямоугольника: %d (x=%d y=%d)\n", x*y, x, y) //
}

func circle() {
	var sq float64
	fmt.Println("Действие: расчет диаметра круга и длины окружности")
	fmt.Print("Введите площадь круга: ")
	n, err := fmt.Scanln(&sq)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	if n == 0 {
		fmt.Println("Ошибка: не введено значение площади круга")
		return
	}

	radius := math.Sqrt(sq / math.Pi)
	fmt.Printf("Диаметр круга: %f\n", radius*2)
	fmt.Printf("Длина окруности: %f\n", 2*math.Pi*radius)
}

func digits() {
	var number int

	fmt.Println("Действие: разряды числа")
	fmt.Print("Введите число от 0 до 999: ")
	n, err := fmt.Scanln(&number)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	if n == 0 {
		fmt.Println("Ошибка: не введено число")
		return
	}

	if number > 999 {
		fmt.Println("Ошибка: введено слишком большое число")
		return
	}

	hundreds := number / 100
	fmt.Printf("Сотни: %d\n", hundreds)
	number = number % 100

	dozens := number / 10
	fmt.Printf("Десятки: %d\n", dozens)
	fmt.Printf("Единицы: %d\n", number%10)
}
