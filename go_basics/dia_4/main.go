package main

import (
	"errors"
	"fmt"
)

type ErrorPersonalizado struct {
	msg string
}

type ErrSalarioMenorDiezMil struct {
}

func (e *ErrorPersonalizado) Error() string {
	return e.msg
}

func (e *ErrSalarioMenorDiezMil) Error() string {
	return "el salario es menor a 10.000"
}

func revisarSalario(salario int) error {
	if salario < 150000000 {
		return &ErrorPersonalizado{"el salario ingresado no alcanza el mínimo imponible"}
	} else {
		fmt.Println("Debe pagar impuesto")
	}
	return nil
}

func aux() {

	var salario int = 180000000000

	// ========= ejercicio 1 =========
	err := revisarSalario(salario)

	if err != nil {
		panic(err)
	}

	// ========= ejercicio 2 =========
	if salario > 10000 {
		err = &ErrSalarioMenorDiezMil{}
		var expectedError ErrSalarioMenorDiezMil
		isSameTypeError := errors.Is(err, &expectedError)
		if isSameTypeError {
			fmt.Println("Los errores son del mismo ")
		} else {
			fmt.Println("Los errores no son del mismo ")
		}
	}
	// ========= ejercicio 3 =========
	if salario > 10000 {
		var expectedError = errors.New("el salario es menor a 10.000")
		err = expectedError
		isSameTypeError := errors.Is(err, expectedError)
		if isSameTypeError {
			fmt.Println("Los errores son del mismo ")
		} else {
			fmt.Println("Los errores no son del mismo ")
		}
	}
	// ========= ejercicio 4 =========
	if salario > 10000 {
		var expectedError ErrSalarioMenorDiezMil
		err = fmt.Errorf("salario mayor a 10000 (%d): %w", salario, &expectedError)
		isSameTypeError := errors.Is(err, &expectedError)
		if isSameTypeError {
			fmt.Println("Los errores son del mismo ")
		} else {
			fmt.Println("Los errores no son del mismo ")
		}
	}
	// ========= ejercicio 5 =========
	total, err := calcularSalario(10, 60)
	if err != nil {
		panic(err)
	}
	fmt.Println("total de horas trabajadas", total)
}

func calcularSalario(horas, valorHora float64) (float64, error) {
	total := horas * valorHora
	if total > 150000000 {
		total *= 0.9
	}
	if horas < 80 {
		return 0, errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	return total, nil
}
