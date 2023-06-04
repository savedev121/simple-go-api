package api

import "github.com/labstack/echo/v4"

func (api *API) RegisterRoutes(e *echo.Echo) {
	prefix := e.Group("api")
	v1 := prefix.Group("v1")
	users := v1.Group("users")
	users.POST("/", api.RegisterUser)

}
