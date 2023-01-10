package main

import (
	"fmt"
	"os"
)

type Client struct {
	legajo    string
	nombre    string
	direccion string
	dni       int
	telefono  int
}

var clientes []Client

func main() {
	// ======== ejercicio 1 y 2 ========
	f, err := os.Open("customers.txt")
	defer f.Close()
	if err != nil {
		err = fmt.Errorf("el archivo indicado no fue encontrado o está dañado: %w", err)
		panic(err)
	}

	// ======== ejercicio 3 ========
	defer fmt.Println("Fin de la ejecucion")
	defer func() {
		err := recover()
		fmt.Println("se detectaron los siguientes errores: ", err)
	}()
	defer func() {
		err := recover()
		if err != nil {
			for _, client := range clientes {
				verificarDatos(client.legajo, client.nombre, client.dni, client.telefono, client.direccion)
			}
		}
	}()
	for _, client := range clientes {
		verifcarClient(client.dni, clientes)
		verificarDatos(client.legajo, client.nombre, client.dni, client.telefono, client.direccion)
	}
}

func verificarDatos(legajo, nombre string, dni, telefono int, domicilio string) error {
	if legajo == "" {
		panic("legajo esta vacio")
	}

	if nombre == "" {
		panic("nombre esta vacio")
	}

	if dni == 0 {
		panic("dni es zero")
	}
	if telefono == 0 {
		panic("telefono es zero")
	}
	if domicilio == "" {
		panic("domicilio esta vacio")
	}
	return nil
}

func verifcarClient(dni int, dniClients []Client) {
	for _, tmp := range dniClients {
		if tmp.dni == dni {
			panic("este cliente ya ha sido registrado anteriormente")
		}
	}
}
