package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	actions := []action{
		{"Площадь прямоугольника", square},
		{"Диаметр круга и длина окружности", circle},
		{"Разряды числа", digits},
		{"Простые числа", primeNumbers},
	}
	menu(actions)
}

// action - действие для меню
type action struct {
	name string
	do   func()
}

// menu организует выбор дейстия из меню
func menu(actions []action) {
	var choice int

	for {
		choice = 0

		fmt.Println("Выберите действие:")
		for c, a := range actions {
			fmt.Printf("%d: %s\n", c+1, a.name)
		}
		fmt.Print("Введите номер действия (0 - выход): ")
		n, err := fmt.Scan(&choice)
		if err != io.EOF {
			if err != nil {
				fmt.Printf("\nОшибка: %v\n", err)
				continue
			}
		}
		if n == 0 {
			fmt.Println("\nОшибка: не выбрано действие")
			return
		}

		if choice == 0 {
			break
		}
		if choice > len(actions) {
			fmt.Printf("\nОшибка: действие с номером %d отсутствует\n", choice)
		} else {
			actions[choice-1].do()
		}
	}
}

// square - вычисляет площать прямоугольника
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

// circle - по площади круга вычисляет диаметр круга и длину окружности
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

// digits - выводит количество сотен, десятков и единиц числа
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

// primeNumbers - выводит простые числа
func primeNumbers() {
	var maxNumber int

	fmt.Println("Действие: определение простых чисел")
	fmt.Print("Введите верхнюю границу диапазона: ")
	n, err := fmt.Scanln(&maxNumber)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	if n == 0 {
		fmt.Println("Ошибка: не введено число")
		return
	}
	if maxNumber > 1000 {
		fmt.Println("Предупреждение: введено число больше 1000")
		return
	}

	// инициализация массива
	nums := make([]bool, maxNumber+1)
	for i := range nums {
		nums[i] = true
	}
	fmt.Println(nums)

	// вычеркивание непростых чисел
	for i := 2; i*i <= maxNumber; i++ {
		if nums[i] {
			for j := i * i; j <= maxNumber; j += i {
				fmt.Println("j = ", j)
				nums[j] = false
			}
		}
	}

	// вывод простых чисел
	firstNum := true
	for i := 2; i <= maxNumber; i++ {
		if nums[i] {
			if firstNum {
				firstNum = false
			} else {
				fmt.Print(", ")
			}
			fmt.Print(i)
		}
	}
	fmt.Println()
}
