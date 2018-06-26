package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inferencia de tipo
	i += 3 // incremento i = i + 3
	i -= 3 // decremento i = i - 3
	i /= 2 // divisao i = i / 2
	i *= 2 // multiplicacao  i = i * 2
	i %= 2 // modulo i = i % 2 //resto da divisao

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y))

	x, y = y, x
	fmt.Println(x, y)

}
