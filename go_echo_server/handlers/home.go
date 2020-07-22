package handlers


import (

	"github.com/labstack/echo"

	"net/http"
)



func Welcome() echo.HandlerFunc {

    return func (c echo.Context) error {
                    return c.JSON(http.StatusOK, "tasks")
			}
			
}