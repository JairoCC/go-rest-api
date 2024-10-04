package handler

import (
	"github.com/JairoCC/go-rest-api/middleware"
	"github.com/labstack/echo/v4"
)

func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("v1/persons")
	person.Use(middleware.Authentication)
	person.POST("", h.create)
	person.PUT("/:id", h.update)
	person.DELETE("/:id", h.delete)
	person.GET("", h.getAll)
	person.GET("/:id", h.getByID)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)
}
