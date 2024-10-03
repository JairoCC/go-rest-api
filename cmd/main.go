package main

import (
	"log"
	"net/http"

	"github.com/JairoCC/go-rest-api/handler"
	"github.com/JairoCC/go-rest-api/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()
	handler.RoutePerson(mux, &store)
	log.Println("servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("error en el servidor: %v\n", err)
	}
}
