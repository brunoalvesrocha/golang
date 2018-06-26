package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//operadores relacionais sempre retornam true or false
	fmt.Println("Strings:", "banana" == "banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 < 2)
	fmt.Println("<=", 3 < 2)
	fmt.Println(">=", 3 < 2)

	d1 := time.Unix(0, 0) //cria uma data = 0,0 referencia da data inicial do calendario
	d2 := time.Unix(0, 0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct { // struct = cria um objeto do tipo Pessoa (em java uma classe) e atribui o valor Joao
		Nome string
	}
	p1 := Pessoa{"Joao"}
	p2 := Pessoa{"Joao"}
	fmt.Println("Pessoas:", p1 == p2)
	fmt.Println(p1)
	fmt.Println(reflect.TypeOf(p1))
}
