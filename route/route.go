package route

import (
	"bowery-golang_testing/controller"

	"github.com/labstack/echo"
)

func SetupRoute(e *echo.Echo) {
	e.POST("/createUser", controller.CreateUser)
	e.GET("/getAllUsers", controller.GetAllUsers)
	e.POST("/chargeUser", controller.ChargeUser)
	e.File("/favicon.ico", "gopher.png")
}
