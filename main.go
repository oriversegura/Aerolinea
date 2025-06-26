package main

import (
	aerolinea "aerolinea/cmd"
	"fmt"
	"log"
)

var seleccion int
var archivo = "pasajeros.json"

func main() {
	// menuPrincipal : Esto imprime el menu principal y acepta un valor entero
	// y retorna un entero con el numero de la seleccion.
	var menu [4]string
	menu[0] = "Registro de Pasajero"
	menu[1] = "Compra de Vuelo"
	menu[2] = "Cancelacion de Vuelo"
	menu[3] = "Disponibilidad de Asientos"

	for {

		// Mostrar el menu
		fmt.Println("Seleccione una opcion:")
		for i, opcion := range menu {
			fmt.Printf("%d. %s\n", i+1, opcion)
		}

		// Preguntar la opcion al usuario
		fmt.Println("Ingrese la opcion deseada: ")
		_, err := fmt.Scanln(&seleccion)
		if err != nil {
			log.Fatal(err)
		}
		if seleccion < 1 || seleccion > len(menu) {
			fmt.Println("Seleccion invalida, intente nuevamente \n")
		} else {
			break
		}
	}
	var seleccion int

	switch seleccion {
	case 1:
		fmt.Println(menu[0])
		err := aerolinea.RegistrarPasajero(archivo)
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		fmt.Println(menu[1])
		aerolinea.ComprarVuelo()
	case 3:
		fmt.Println(menu[2])
	case 4:
		fmt.Println(menu[3])
	default:
		break
	}
}
