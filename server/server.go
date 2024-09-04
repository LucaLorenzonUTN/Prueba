package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lib"
)

func main() {
	http.HandleFunc("POST /helloworld", HelloWorld)
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	var request lib.BodyRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var respBody lib.BodyResponse = lib.BodyResponse{
		Mensaje: fmt.Sprintf("Hola %s! y el pedido es: %s ", request.Name, request.Pedido),
	}

	respuesta, err := json.Marshal(respBody)
	if err != nil {
		http.Error(w, "Error al codificar los datos como JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)

	fmt.Printf("Me hablo %s y me pidio %s \n", request.Name, request.Pedido)
}
