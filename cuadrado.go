package figuras

// clase cuadrado
// propiedades y metodos
type Cuadrado struct {
	Lado float64
}

func (cua *Cuadrado) area() float64 {
	return cua.Lado * cua.Lado
}

func (cua *Cuadrado) perimetro() float64 {
	return 4 * cua.Lado
}
