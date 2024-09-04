package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"lib"
)

func main() {

	nombre, funcion := pedirMsgConsola()
	body := armarMensaje(nombre, funcion)

	cliente := &http.Client{}
	url := fmt.Sprintf("http://localhost:8080/helloworld")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	respuesta, err := cliente.Do(req)
	if err != nil {
		return
	}

	// Verificar el código de estado de la respuesta
	if respuesta.StatusCode != http.StatusOK {
		return
	}

	var response lib.BodyResponse
	err = json.NewDecoder(respuesta.Body).Decode(&response)
	if err != nil {
		return
	}

	fmt.Println(response.Mensaje)
}

func pedirMsgConsola() (nombre string, funcion string) {

	fmt.Println("Ingresa Nombre y funcion serparado por un espacio")
	fmt.Scan(&nombre, &funcion)
	return nombre, funcion

}

func armarMensaje(Nombre string, Funcion string) []byte {

	body, err := json.Marshal(lib.BodyRequest{
		Name:   Nombre,
		Pedido: Funcion,
	})
	if err != nil {
		panic("Error al crear msg")
	}

	return body

}
