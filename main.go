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
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		num, err := strconv.ParseInt(scan.Text(), 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		arr = append(arr, int(num))
	}
	fmt.Println("Unsorted: ", arr)
	quickSort(arr)
	fmt.Println("Sorted: ", arr)
}

func quickSort(arr []int) {
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
	quickSort(arr[:rightInd+1])
	// .. правой части (в ней элементы больше опорного)
	quickSort(arr[leftInd:])
}
