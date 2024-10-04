package middleware

import (
	"net/http"

	"github.com/JairoCC/go-rest-api/authorization"
	"github.com/labstack/echo/v4"
)

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Unauthorized"})
		}

		return f(c)
	}
}
