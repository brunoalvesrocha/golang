package main

import (
	"fmt"
)

// Não tem operador ternário
func obtemResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obtemResultado(6.2))
}
