package controllers

import (
	"flycode.go/abacusf/app/controllers/dashboard"
	"flycode.go/abacusf/app/controllers/front"
	"flycode.go/abacusf/app/tools/debug"
	"github.com/labstack/echo"
)

func Process(e *echo.Echo) {
	e.HTTPErrorHandler = customHTTPErrorHandler
	front.Front(e)
	dashboard.Dashboard(e)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	debug.Error("%v !!!! Main Error", err)
}