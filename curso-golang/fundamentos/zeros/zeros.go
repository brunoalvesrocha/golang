package main

import (
	"fmt"
)

func main() {

	var a int
	var b float64
	var c bool
	var d string // defaul vazia quando nao eh iniciada com valor
	var e *int   // ponteiro por default <nil> quando nao eh iniciada

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
