package main

import "fmt"

type action struct {
	name string
	do   func()
}

func menu(actions []action) {

	for {
		fmt.Println("Действия:")
		for c, a := range actions {
			fmt.Printf("%d: %s\n", c+1, a.name)
		}
		fmt.Print("Введите номер действия (0 - выход): ")
		choose := 0
		_, err := fmt.Scan(&choose)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}
		if choose == 0 {
			break
		}

		if choose > len(actions) {
			fmt.Printf("Ошибка: действие с номером %d отсутствует\n", choose)
		} else {
			actions[choose-1].do()
		}
	}

}

func main() {
	actions := []action{
		{"Площадь прямоугольника", square},
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
