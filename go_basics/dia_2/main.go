package main

import "errors"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func ejercicio_1(salario float64) float64 {
	total := 0.0
	if salario > 50000.0 {
		total += salario * 0.17
	}
	if salario > 150000.0 {
		total += (salario - 150000.0) * 0.1
	}
	return total
}

func ejercicio_2(notas ...int) (int, error) {
	n := len(notas)
	if n == 0 {
		return 0, errors.New("Debe haber minimo una nota")
	}

	total := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Las notas no pueden ser negativas")
		}
		total += nota
	}
	promedio := total / n

	return promedio, nil
}

func ejericio_3(minutos int, categoria string) int {
	horas := minutos / 60
	var pago int
	switch categoria {
	case "A":
		pago = horas * 3000
		pago += pago / 2
	case "B":
		pago = horas * 1500
		pago += pago / 5
	case "C":
		pago = horas * 1000
	}
	return pago
}

func ejercicio_4(tipo string) (func(...int) (int, error), error) {
	if tipo == minimum {
		return func(notas ...int) (int, error) {
			if len(notas) == 0 {
				return 0, errors.New("No se ingresaron valores")
			}
			respuesta := notas[0]
			for _, val := range notas {
				if val < respuesta {
					respuesta = val
				}
			}
			return respuesta, nil
		}, nil
	} else if tipo == maximum {
		return func(notas ...int) (int, error) {
			if len(notas) == 0 {
				return 0, errors.New("No se ingresaron valores")
			}
			respuesta := notas[0]
			for _, val := range notas {
				if val > respuesta {
					respuesta = val
				}

			}
			return respuesta, nil
		}, nil
	} else if tipo == average {
		return func(notas ...int) (int, error) {
			if len(notas) == 0 {
				return 0, errors.New("No se ingresaron valores")
			}
			total := 0
			n := len(notas)
			for _, val := range notas {
				total += val
			}
			return total / n, nil
		}, nil
	} else {
		return nil, errors.New("Tipo de operacion no soportada")
	}
}

func ejercicio_5(tipo string) (func(float64) float64, error) {
	var multiplicador float64
	if tipo == cat {
	} else if tipo == dog {
		multiplicador = 5
	} else if tipo == hamster {

		multiplicador = 10
	} else if tipo == hamster {

		multiplicador = 0.25
	} else if tipo == tarantula {

		multiplicador = 0.15
	} else {
		return nil, errors.New("Animal no soportado")
	}
	return func(cantidad float64) float64 {
		return cantidad * multiplicador
	}, nil

}

func main() {
}
