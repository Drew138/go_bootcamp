package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjercicio1_SalarioMenorA50000(t *testing.T) {
	resultadoEsperado := 0.0

	resultado := ejercicio_1(4000)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestEjercicio1_SalarioMayorA50000(t *testing.T) {
	resultadoEsperado := 8670.0

	resultado := ejercicio_1(51000)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestEjercicio1_SalarioMayorA150000(t *testing.T) {
	resultadoEsperado := 25770.000000000004

	resultado := ejercicio_1(151000)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestEjercicio2_ZeroNotas(t *testing.T) {
	expectedError := errors.New("Debe haber minimo una nota")
	expectedResult := 0
	result, err := ejercicio_2()
	assert.Equal(t, expectedResult, result)
	assert.EqualError(t, expectedError, err.Error())
}

func TestEjercicio2_NotaNegativas(t *testing.T) {
	expectedError := errors.New("Las notas no pueden ser negativas")
	expectedResult := 0
	result, err := ejercicio_2(1, 2, 3, -5)
	assert.Equal(t, expectedResult, result)
	assert.EqualError(t, expectedError, err.Error())
}

func TestEjercicio2_Success(t *testing.T) {
	expectedResult := 2
	result, err := ejercicio_2(1, 2, 3, 5)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, nil, err)
}

func TestEjercicio3_CasoA(t *testing.T) {
	expectedResult := 4500
	var caso = "A"
	minutos := 60
	result := ejericio_3(minutos, caso)
	assert.Equal(t, expectedResult, result)
}

func TestEjercicio3_CasoB(t *testing.T) {

	expectedResult := 1800
	var caso = "B"
	minutos := 60
	result := ejericio_3(minutos, caso)
	assert.Equal(t, expectedResult, result)
}

func TestEjercicio3_CasoC(t *testing.T) {
	expectedResult := 1000
	var caso = "C"
	minutos := 60
	result := ejericio_3(minutos, caso)
	assert.Equal(t, expectedResult, result)
}

func TestEjercicio3_CasoDefault(t *testing.T) {
	expectedResult := 0
	var caso = ""
	minutos := 60
	result := ejericio_3(minutos, caso)
	assert.Equal(t, expectedResult, result)
}

func TestEjercicio4_MinimoSinValores(t *testing.T) {

	expectedResult := 0
	var expectedError1 error
	expectedError2 := errors.New("No se ingresaron valores")

	operacion, err1 := ejercicio_4(minimum)
	result, err2 := operacion()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError1, err1)
	assert.EqualError(t, expectedError2, err2.Error())

}

func TestEjercicio4_MinimoConValores(t *testing.T) {

	expectedResult := 1
	var expectedError2, expectedError1 error

	operacion, err1 := ejercicio_4(minimum)
	result, err2 := operacion(1, 2, 3)

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError1, err1)
	assert.Equal(t, expectedError2, err2)

}

func TestEjercicio4_MaximoSinValores(t *testing.T) {

	expectedResult := 0
	var expectedError1 error
	expectedError2 := errors.New("No se ingresaron valores")

	operacion, err1 := ejercicio_4(maximum)
	result, err2 := operacion()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError1, err1)
	assert.EqualError(t, expectedError2, err2.Error())

}

func TestEjercicio4_MaximoConValores(t *testing.T) {

	expectedResult := 3
	var expectedError2, expectedError1 error

	operacion, err1 := ejercicio_4(maximum)
	result, err2 := operacion(1, 2, 3)

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError1, err1)
	assert.Equal(t, expectedError2, err2)

}

func TestEjercicio4_PromedioSinValores(t *testing.T) {

	expectedResult := 0
	var expectedError1 error
	expectedError2 := errors.New("No se ingresaron valores")

	operacion, err1 := ejercicio_4(average)
	result, err2 := operacion()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError1, err1)
	assert.EqualError(t, expectedError2, err2.Error())

}

func TestEjercicio4_PromedioConValores(t *testing.T) {

	expectedResult := 2
	var expectedError2, expectedError1 error

	operacion, err1 := ejercicio_4(average)
	result, err2 := operacion(1, 2, 3)

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError1, err1)
	assert.Equal(t, expectedError2, err2)
}
