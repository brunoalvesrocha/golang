package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado ... pega o valor 97 da tabela ASCII
	fmt.Println("Teste " + string(97))

	// int para String
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123") // atoi devolve um inteiro e um erro caso ocorra _ => ignora o valor da variavel
	fmt.Println(num - 122)

	// string para booleano
	b, _ := strconv.ParseBool("T")
	if b {
		fmt.Println("Verdadeiro")
	}
}
