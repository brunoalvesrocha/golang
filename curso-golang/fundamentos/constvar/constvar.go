package main

import (
	format "fmt"
	mt "math" //pode-se atribuir um label ao import
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //float64 inferido automaticamente -> padrão para ponto flutuante

	/*forma reduzida para declarar variaveis e inicializar*/
	area := PI * mt.Pow(raio, 2)
	format.Println("A área da circunferencia é: ", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)

	format.Println(a, b, c, d)

	var e, f bool = true, false //declara as variaveis e f e atribui booleanos repectivamente true false
	format.Println(e, f)

	g, h, i := 2, false, "epa!" //declara e inicializa variaveis g h i com valores 2 false epa (fortemente tipados, não anteram o tipo ao longo do programa)
	format.Println(g, h, i)

}
