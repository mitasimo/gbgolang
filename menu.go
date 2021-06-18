package main

import (
	"fmt"
	"io"
)

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
