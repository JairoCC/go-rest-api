package handler

import (
	"net/http"

	"github.com/JairoCC/go-rest-api/authorization"
	"github.com/JairoCC/go-rest-api/model"
	"github.com/labstack/echo/v4"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {

	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "object structure is incorrect", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if !isLogingValid(&data) {
		response := newResponse(Error, "user and/or password not valid", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "token could not be generated", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	dataToken := map[string]string{"token": token}
	response := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, response)
}

func isLogingValid(data *model.Login) bool {
	return data.Email == "contact@jhc.com" && data.Password == "12345"
}
