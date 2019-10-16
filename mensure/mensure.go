package mensure

import (
	"fmt"
	"go_interface_exemplo_tutorial/geometry"
)

/*
	Posso usar esta interface aqui ou inportar o pacote
*/

// Geometry - 
//type Geometry interface {
//	Area() float64/
//	Perim() float64
//}

//Mensure  - business rule
func Mensure(g geometry.Geometry) {

	area := g.Area()
	perim := g.Perim()
	if area > 100 {
		fmt.Println("Area limit exceeded!")
	} else {
		fmt.Println(g)
		fmt.Println(area)
		fmt.Println(perim)
	}
}
