package dashboard

import (
	"github.com/labstack/echo"
)

func Dashboard(e *echo.Echo) {
	e.GET("dashboard/login", LoginController)
	e.GET("dashboard/login/", LoginController)
	e.POST("dashboard/login", PLoginController)
	e.POST("dashboard/login/", PLoginController)
	e.GET("dashboard/disconnection", DisconnectionController)
	e.GET("dashboard/disconnection/", DisconnectionController)

	e.GET("dashboard", HomeController)
	e.GET("dashboard/", HomeController)
}
