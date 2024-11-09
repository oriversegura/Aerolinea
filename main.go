package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Pasajero struct {
	ID       int    `json:"ID" validate:"required"`
	Nombre   string `json:"Nombre" validate:"required"`
	Apellido string `json:"Apellido" validate:"required"`
	Edad     int    `json:"Edad" validate:"gte=0,lte=120"`
	Destino  string `json:"Destino" validate:"required"`
}

var (
	pasajero  Pasajero
	archivo   = "pasajeros.json"
	continuar = true
	seleccion = 0
)

func main() {
	// menuPrincipal : Esto imprime el menu principal y acepta un valor entero
	// y retorna un entero con el numero de la seleccion.
	var menu [4]string
	menu[0] = "Registro de Pasajero"
	menu[1] = "Compra de Vuelo"
	menu[2] = "Cancelacion de Vuelo"
	menu[3] = "Disponibilidad de Asientos"

	for continuar {
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
			fmt.Println("Seleccion Invalida, intente nuevamente")
			fmt.Println("")
			continuar = true
		} else {
			continuar = false
		}
	}

	switch seleccion {
	case 1:
		fmt.Println(menu[0])
		err := registrarPasajero(archivo)
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		fmt.Println(menu[1])

	case 3:
		fmt.Println(menu[2])
	case 4:
		fmt.Println(menu[3])
	default:
		break
	}
}

func registrarPasajero(archivo string) error {
	var pasajeros []Pasajero

	err := revisarJSON(archivo)
	if err != nil {
		return err
	}

	file, err := os.ReadFile(archivo)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(file), &pasajeros)

	nuevoPasajero := Pasajero{}

	//Primero preguntamos el ID
	fmt.Println("Inserte el ID de pasajero: ")
	fmt.Scanf("%d", &nuevoPasajero.ID)

	//Nombre, Apellido y Destino que sean multiple.
	fmt.Println("Inserte el nombre del pasajero (enter para finalizar): ")
	fmt.Scanln(&nuevoPasajero.Nombre)
	fmt.Println("Inserte el apellido del pasajero (enter para finalizar): ")
	fmt.Scanln(&nuevoPasajero.Apellido)
	fmt.Println("Inserte la edad del pasajero: ")
	fmt.Scanf("%d", &nuevoPasajero.Edad)
	fmt.Println("Inserte el nombre del pasajero (enter para finalizar): ")
	fmt.Scan(&nuevoPasajero.Destino)

	pasajeros = append(pasajeros, nuevoPasajero)

	pasajeroJSON, err := json.MarshalIndent(pasajeros, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(pasajeroJSON))

	err = os.WriteFile(archivo, pasajeroJSON, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Pasajero registrado exitosamente!")

	return nil
}

func revisarJSON(archivo string) error {
	_, err := os.Stat(archivo)
	if os.IsNotExist(err) {
		_, err := os.Create(archivo)
		if err != nil {
			return err
		}
	}
	return nil
}
