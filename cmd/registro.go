package aerolinea

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

func RegistrarPasajero(archivo string) error {
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

	// Primero preguntamos el ID
	fmt.Println("Inserte el ID de pasajero: ")
	fmt.Scanf("%d", &nuevoPasajero.ID)

	// Nombre, Apellido y Destino que sean multiple.
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
