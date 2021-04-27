package main

import (
	"fmt"

	"github.com/mitasimo/gbgolang/fibonachi"
)

func main() {

	const (
		n  = 500
		ln = 30
	)

	fmt.Printf("FibVars(%d)\t= %d\n", n, fibonachi.FibVars(n))
	fmt.Printf("FibCache(%d)\t= %d\n", n, fibonachi.FibCache(n))
	fmt.Printf("FibRecur(%d)\t= %d\n", ln, fibonachi.FibRecur(ln))

}
