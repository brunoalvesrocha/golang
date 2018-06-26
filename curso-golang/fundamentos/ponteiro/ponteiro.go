package main

import (
	"fmt"
)

func main() {
	i := 1 //64bits = 8bytes

	var p *int = nil //cria um ponteiro do tipo inteiro e armazena 64bits de memória.

	p = &i //& -> pega o endereco da variavel I e armazena em P
	*p++   // *p desreferencia o endereco da memoria e pega o valor do conteudo I
	i++
	fmt.Println(p, *p, i, &i)

	//go nao tem aritimetica de ponteiros (referencia de memória)
	//p++
}
