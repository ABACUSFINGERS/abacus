package controllers

import (
	"./dashboard"
	"./front"
	"../tools/debug"
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