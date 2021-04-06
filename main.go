package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var arr []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		arr = append(arr, int(num))
	}
	fmt.Println("Unsorted: ", arr)
	SortInsert(arr)
	fmt.Println("Sorted: ", arr)
}

func SortQuick(arr []int) {
	if len(arr) < 2 {
		return
	}

	// индексы границы массива
	leftInd, rightInd := 0, len(arr)-1
	// опорный элемент
	pivot := arr[(leftInd+rightInd)/2]

	for {
		// искать сначала элемент, который меньше опорного
		for ; arr[leftInd] < pivot; leftInd++ {
		}
		// искать с конца элемент, который больше опорного
		for ; arr[rightInd] > pivot; rightInd-- {
		}
		// если индекс меньшего не больше индекса большего, обменять значения местами
		if leftInd <= rightInd {
			arr[leftInd], arr[rightInd] = arr[rightInd], arr[leftInd]
			leftInd++
			rightInd--
		}
		if leftInd > rightInd {
			break
		}
	}
	// рекурсивно вызвать для левой части (в ней элементы меньше опорного)
	SortQuick(arr[:rightInd+1])
	// .. правой части (в ней элементы больше опорного)
	SortQuick(arr[leftInd:])
}

func SortInsert(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
}
