package main

import (
	"fmt"

	"github.com/mitasimo/gbgolang/fibonachi"
)

func main() {

	const (
		n  = 10000
		ln = 30
	)

	fmt.Printf("FibVars(%d)\t= %d\n", n, fibonachi.FibVars(n))
	fmt.Printf("FibSlice(%d)\t= %d\n", n, fibonachi.FibSlice(n))
	fmt.Printf("FibMap(%d)\t= %d\n", n, fibonachi.FibMap(n))
	fmt.Printf("FibRecur(%d)\t= %d\n", ln, fibonachi.FibRecur(ln))

}
