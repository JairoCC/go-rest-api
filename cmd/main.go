package main

import (
	"log"

	"github.com/JairoCC/go-rest-api/authorization"
	"github.com/JairoCC/go-rest-api/handler"
	"github.com/JairoCC/go-rest-api/storage"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("certifacates cannot be loaded: %v", err)
	}
	store := storage.NewMemory()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// mux := http.NewServeMux()
	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)
	log.Println("servidor iniciado en el puerto 8080")
	err = e.Start(":8080")
	if err != nil {
		log.Printf("error en el servidor: %v\n", err)
	}
}
