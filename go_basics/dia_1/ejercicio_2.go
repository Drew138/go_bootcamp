package main

import "fmt"

func ejericio_1() {
	var palabra string
	fmt.Scanln(&palabra)
	caracteres := [26]int{}
	for _, character := range palabra {
		caracteres[int(character)-int('a')]++
	}
	for i, cantidad := range caracteres {
		fmt.Printf("Existen %d caracteres de %c\n", cantidad, int('a')+i)
	}
}

func ejericio_2() {
	var (
		edad       int
		empleado   bool
		sueldo     float64
		antiguedad int
	)

	fmt.Scan(&edad, &empleado, &sueldo, &antiguedad)
	if !empleado {
		fmt.Println("No esta empleado, no se le otorga prestamo")
	}

	if edad < 22 {
		fmt.Println("Tiene menos de 22 a;os, no se le otorga prestamo")
	}

}

func ejercicio_4() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count = 0
	for _, val := range employees {
		if val >= 21 {
			count++
		}
	}
	fmt.Printf("Numero de empleados mayores de 21: %d\n", count)

}

// func main() {
// 	ejericio_1()
// 	ejericio_2()
// }
