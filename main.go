package main

import (
	"fmt"
	"go_interface_exemplo_tutorial/mensure"
)

func main() {

	/*
		Essa parte equivale a chamada do construtor da classe, que,
		na verdade estamos apenas populando uma struct
	*/
	r := Rect{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	/*
		Aqui uso a interface para calcular a area e o perimetro, não importa o objeto passado,
		pode ser um retangulo ou um círculo desde que cumpra o contrato com a interface Geometry
	*/
	mensure.Mensure(r)
	mensure.Mensure(c)

	
	/*
		Tenho acesso aos metodos diretamente mas no caso, não obedece as regras de negócio implementadas
		em mensure
	*/
	fmt.Println(r)
	fmt.Println(r.Area())
	fmt.Println(r.Perim())
	fmt.Println(c)
	fmt.Println(c.Area())
	fmt.Println(c.Perim())

}
