package main

import "fmt"

func RunFibonachi() {
	var num int

	fmt.Print("Введите номер числа: ")
	fmt.Scanln(&num)
	fmt.Printf("Фибоначчи(%d) = %d\n", num, fibonachiVar(num))
}

func fibonachiVar(n int) int {
	var fib int

	if n <= 1 {
		return n
	}

	n0, n1 := 0, 1
	for i := 2; i <= n; i++ {
		fib = n0 + n1
		n0, n1 = n1, fib
	}

	return fib
}

func fibonachiMap(n int) int {
	if n <= 1 {
		return n
	}

	fib := make(map[int]int)
	fib[0], fib[1] = 0, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}

	return fib[n]

}

func fibonachiRecur(n int) int {
	if n <= 1 {
		return n
	}

	return fibonachiRecur(n-2) + fibonachiRecur(n-1)
}
