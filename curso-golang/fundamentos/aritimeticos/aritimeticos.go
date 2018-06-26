package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtracao =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("multiplicacao =>", a*b)
	fmt.Println("modulo =>", a%b) //pega o resto da divisao entre a e b

	//bitwise => quebra o valor em bits e efetua uma operacao
	fmt.Println("E =>", a&b)  // 11 & 10 = 2
	fmt.Println("Ou =>", a|b) // 11 | 10 = 3

	c := 3.0
	d := 2.0

	//outras operacoes utlizando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciacao =>", math.Pow(c, d))
}
