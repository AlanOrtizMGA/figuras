package figuras

import (
	"fmt"
)

// interfaz
// los metodos pueden ser privados
type Geometria interface {
	area() float64
	perimetro() float64
}

// orquestador es publico
func Medidas(g Geometria) {
	fmt.Println("medidas", g)
	fmt.Println("Area", g.area())
	fmt.Println("perimetro", g.perimetro())
}
