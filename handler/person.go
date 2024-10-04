package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JairoCC/go-rest-api/model"
	"github.com/labstack/echo/v4"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {

	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "object structure is incorrect", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "an error occurred when creating person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "person created successfully", data)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID must be a positive integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "object structure is incorrect", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Update(ID, &data)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "person ID does not exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "an error occurred when updating a person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID must be a positive integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "person ID does not exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "an error occurred when deleting a person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := model.Person{}
	err = c.Bind(&data)
	response := newResponse(Message, "Ok", err)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getAll(c echo.Context) error {
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "an error occurred when trying to get all persons", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID must be a positive integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "person ID does not exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err != nil {
		response := newResponse(Error, "an error occurred when trying to get a persons", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}
