package main

import (

	"github.com/labstack/echo"
	"abacus/go_echo_server/handlers"
	"github.com/labstack/echo/middleware"
)



	
func main() {

	// Echo instance
	e := echo.New()
  
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	// Routes
	e.GET("/", handlers.Welcome())
	
  
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
  }
