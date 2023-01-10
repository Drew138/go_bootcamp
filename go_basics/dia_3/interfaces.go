package main

import "fmt"

const (
	pequeno = "pequeno"
	mediano = "mediano"
	grande  = "grande"
)

// ========= ejercicio 1 =========
type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (e *Estudiante) detalle() {
	fmt.Printf("Nombre: %s\n", e.Nombre)
	fmt.Printf("Apellido: %s\n", e.Apellido)
	fmt.Printf("DNI: %d\n", e.DNI)
	fmt.Printf("Fecha: %s\n", e.Fecha)
}

// ========= ejercicio 2 =========

type ProductoTienda interface {
	Costo() float64
}

type ProductoPequeno struct {
	precio float64
}

type ProductoMediano struct {
	ProductoPequeno
	mantenimiento float64
}

type ProductoGrande struct {
	ProductoMediano
	envio float64
}

func (p *ProductoPequeno) Costo() float64 {
	return p.precio
}

func (p *ProductoMediano) Costo() (total float64) {
	total = p.precio + p.mantenimiento
	return

}

func (p *ProductoGrande) Costo() (total float64) {
	total = p.precio + p.mantenimiento + p.envio
	return
}

func NewProductoTienda(tipoProducto string, args ...float64) ProductoTienda {
	switch tipoProducto {
	case pequeno:
		precio := args[0]
		return &ProductoPequeno{precio}
	case mediano:
		precio := args[0]
		porcentaje := args[1]
		return &ProductoMediano{ProductoPequeno{precio}, precio * porcentaje}
	case grande:
		precio := args[0]
		porcentaje := args[1]
		costoEnvio := args[2]
		return &ProductoGrande{ProductoMediano{ProductoPequeno{precio}, precio * porcentaje}, costoEnvio}
	default:
		return nil
	}
}

func main() {

	productoUno := NewProductoTienda(pequeno, 5000)
	productoDos := NewProductoTienda(mediano, 10000, 0.03)
	productoTres := NewProductoTienda(grande, 15000, 0.06, 2500)

	fmt.Printf("%v\n", productoUno.Costo())
	fmt.Printf("%v\n", productoDos.Costo())
	fmt.Printf("%v\n", productoTres.Costo())

}
