package controllers

import (
	"mindset.go/abacus/app/controllers/dashboard"
	"mindset.go/abacus/app/controllers/front"
	"mindset.go/abacus/app/tools/debug"
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