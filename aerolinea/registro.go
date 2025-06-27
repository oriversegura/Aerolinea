package aerolinea

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Inserte el ID de pasajero: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		return fmt.Errorf("ID invalido: %w", err)
	}
	nuevoPasajero.ID = id

	fmt.Println("Inserte el nombre del pasajero: ")
	nombre, _ := reader.ReadString('\n')
	nuevoPasajero.Nombre = strings.TrimSpace(nombre)

	fmt.Println("Inserte el apellido del pasajero: ")
	apellido, _ := reader.ReadString('\n')
	nuevoPasajero.Apellido = strings.TrimSpace(apellido)

	fmt.Println("Inserte la edad del pasajero: ")
	edadStr, _ := reader.ReadString('\n')
	edad, err := strconv.Atoi(strings.TrimSpace(edadStr))
	if err != nil {
		return fmt.Errorf("edad invalida: %w", err)
	}
	nuevoPasajero.Edad = edad

	fmt.Println("Inserte el destino del pasajero: ")
	destino, _ := reader.ReadString('\n')
	nuevoPasajero.Destino = strings.TrimSpace(destino)

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